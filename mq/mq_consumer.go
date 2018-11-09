package main

import (
	"encoding/json"
	"github.com/cihub/seelog"
	"github.com/streadway/amqp"
	"io/ioutil"
	"nonolive/nonoutils/mq"
	"strconv"
	"strings"
	"time"
)

var NormalMqUrl string
var HJBExchange string
var HJBRoutingKey string
var HJBQueue string
var HJBMqChannel *mq.RabbitMqChannelHelper
var HJBMqConnHelper *mq.RabbitMqConnHelper
var logger seelog.LoggerInterface
var statistic _Statistic

type _Statistic struct {
	StartTime       time.Time
	EndTime         time.Time
	SuccMsg         int
	TotalMsg        int
	SuccRate        float64
	SuccMsgDuration float64 //用于计算一条消息发送消费平均耗时  SuccMsgDuration 除以 SuccMsg
}

type _Msg struct {
	StartTime time.Time `json:"starttime"`
	EndTime   time.Time `json:"endtime"`
	Index     int       `json:"index"`
	TotalMsg  int       `json:"total_msg"`
	IsStart   bool      `json:"is_start"`
	IsEnd     bool      `json:"is_end"`
}

func main() {
	port := 1111
	datas, err := ioutil.ReadFile("./seelog-mq.xml")
	contentStr := string(datas)
	contentStr = strings.Replace(contentStr, "##{port}", strconv.Itoa(port), -1)
	logger, err = seelog.LoggerFromConfigAsString(contentStr)
	if err != nil {
		logger.Error(err)
		return
	}
	NormalMqUrl = "amqp://admin:admin@192.168.16.20:5672/"
	HJBExchange = "hjb_exchange"
	HJBRoutingKey = "hjb"
	HJBQueue = "hjb_queue"
	HJBMqConnHelper = new(mq.RabbitMqConnHelper)
	err = HJBMqConnHelper.Init(NormalMqUrl, logger)
	if err != nil {
		logger.Errorf("Init MqConn fail,err ： %v", err)
		return
	}
	HJBMqChannel = new(mq.RabbitMqChannelHelper)
	err = HJBMqChannel.Init("hjb", HJBMqConnHelper, registerListener, logger)
	if err != nil {
		logger.Error("Init mq channel fail", err)
		return
	}
	DeclareHJBExchange()
	DeclareHJBQueue()
	consumeHJBQueue()
	logger.Info("consumer start")
	for {
		time.Sleep(time.Second * 1)
	}
}

func DeclareHJBExchange() {
	if HJBMqConnHelper == nil {
		logger.Error("nil mq helper")
		return
	}
	ch := HJBMqChannel.EffectChannel
	if ch == nil {
		logger.Error("Declare channel open channel is nil")
		return
	}
	err := ch.ExchangeDeclare(HJBExchange, "direct", true, false, false, false, nil)
	if err != nil {
		logger.Errorf("exchange declare fail,err is %v", err)
	}
}

func DeclareHJBQueue() {
	if HJBMqConnHelper == nil {
		logger.Error("nil mq helper")
		return
	}
	ch := HJBMqChannel.EffectChannel
	if ch == nil {
		logger.Error("open new channel is nil")
		return
	}
	_, err := ch.QueueDeclare(HJBQueue, true, false, false, false, nil)
	if err != nil {
		logger.Errorf("queue decalre fail,err is :%v", err)
		return
	}
	err = ch.QueueBind(HJBQueue, HJBRoutingKey, HJBExchange, false, nil)
	if err != nil {
		logger.Errorf("queue bind is fail", err)
		return
	}
	/*
		_, err = ch.QueueDeclare(HJBQueue2, true, false, false, false, nil)
		if err != nil {
			logger.Errorf("queue decalre fail,err is :%v", err)
			return
		}
		err = ch.QueueBind(HJBQueue2, HJBRoutingKey, HJBExchange, false, nil)
		if err != nil {
			logger.Errorf("queue bind is fail", err)
			return
		}
	*/
}

func consumeHJBQueue() {
	go func() {
		tryTimes := 1
	TryConsumeLabel:
		if tryTimes >= 3 {
			time.Sleep(time.Second)
		}
		if HJBMqChannel == nil {
			logger.Errorf("NewFansMqChannel is nil. skip out.")
			return
		}
		msgChan, err := HJBMqChannel.WrapConsume(HJBQueue, "", false,
			false, false, false, nil)
		if err != nil {
			logger.Errorf("create consumer fail. %v", err)
		} else {
			for msg := range msgChan {
				msg.Ack(false) //消费者需要ack告诉生产者我已经收到了！
				var _msg _Msg
				json.Unmarshal(msg.Body, &_msg)
				_msg.EndTime = time.Now()
				statistic.TotalMsg = _msg.TotalMsg
				statistic.SuccMsg++
				statistic.SuccMsgDuration += _msg.EndTime.Sub(_msg.StartTime).Seconds()
				if _msg.IsStart {
					statistic.StartTime = time.Now()
				}
				//logger.Infof("task:%d,time ns %+v", _msg.Index, _msg.EndTime.Sub(_msg.StartTime).Nanoseconds())
				if _msg.IsEnd {
					statistic.EndTime = time.Now()
					statistic.SuccRate = float64(statistic.SuccMsg) / float64(statistic.TotalMsg)
					//endChan <- statistic
					logger.Infof("总测试条数:%d,成功条数:%d ,成功率:%0.3f,总耗时:%f /秒,一条消息耗时:%f/秒",
						statistic.TotalMsg, statistic.SuccMsg, statistic.SuccRate, statistic.EndTime.Sub(statistic.StartTime).Seconds(),
						float64(statistic.SuccMsgDuration/float64(statistic.SuccMsg)))
					statistic = _Statistic{}
				}
			}
			tryTimes = 0
			logger.Errorf("new fans mq msgChan close.timeTimes is %v", tryTimes)
		}
		tryTimes++
		goto TryConsumeLabel
	}()
}

func registerListener(channel *amqp.Channel) {
	channel.Confirm(false)
	notifyPublishChan := channel.NotifyPublish(make(chan amqp.Confirmation))
	go func() {
		for t := range notifyPublishChan {
			logger.Infof("notifyPublishChan call. %+v", t)
		}
	}()
	cancelChan := channel.NotifyCancel(make(chan string))
	go func() {
		for t := range cancelChan {
			logger.Infof("notifyCancel call. %+v", t)
		}
	}()
	closeChan := channel.NotifyClose(make(chan *amqp.Error))
	go func() {
		for t := range closeChan {
			logger.Infof("notifyClose call. %+v", t)
		}
	}()
	ackChan, nackChan := channel.NotifyConfirm(make(chan uint64), make(chan uint64))
	go func() {
		for t := range ackChan {
			logger.Infof("notifyConfirm call. ack. %+v", t)
		}
	}()
	go func() {
		for t := range nackChan {
			logger.Infof("notifyConfirm call. nack. %+v", t)
		}
	}()
	flowChan := channel.NotifyFlow(make(chan bool))
	go func() {
		for t := range flowChan {
			logger.Infof("notifyFlow call. %+v", t)
		}
	}()
	returnChan := channel.NotifyReturn(make(chan amqp.Return))
	go func() {
		for t := range returnChan {
			logger.Infof("notifyReturn call. %+v", t)
		}
	}()
}
