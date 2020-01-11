package main

import (
	"encoding/json"
	"flag"
	"github.com/cihub/seelog"
	"github.com/streadway/amqp"
	"io/ioutil"
	"nonolive/nonoutils/mq"
	"strconv"
	"strings"
	"sync"
	"time"
)

var NormalMqUrl string
var HJBExchange string
var HJBRoutingKey string
var HJBQueue string
var HJBMqChannel *mq.RabbitMqChannelHelper
var HJBMqChannel2 *mq.RabbitMqChannelHelper
var HJBMqConnHelper *mq.RabbitMqConnHelper
var HJBMqConnHelper2 *mq.RabbitMqConnHelper
var logger seelog.LoggerInterface
var times *int
var statistic _Statistic
var wg sync.WaitGroup

type _Statistic struct {
	StartTime time.Time
	EndTime   time.Time
	TryTimes  int
	SuccTimes int
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
	times = flag.Int("times", 100000, "测试条数")
	flag.Parse()
	port := 1111
	datas, err := ioutil.ReadFile("./seelog-mq2.xml")
	contentStr := string(datas)
	contentStr = strings.Replace(contentStr, "##{port}", strconv.Itoa(port), -1)
	logger, err = seelog.LoggerFromConfigAsString(contentStr)
	if err != nil {
		logger.Error(err)
		return
	}
	statistic.TryTimes = *times
	NormalMqUrl = "amqp://admin:admin@192.168.16.20:5672/"
	HJBExchange = "hjb_exchange"
	HJBRoutingKey = "hjb"
	HJBQueue = "hjb_queue"
	HJBMqConnHelper = new(mq.RabbitMqConnHelper)
	err = HJBMqConnHelper.Init(NormalMqUrl, nil)
	if err != nil {
		logger.Errorf("Init MqConn fail,err ： %v", err)
		return
	}
	HJBMqConnHelper2 = new(mq.RabbitMqConnHelper)
	err = HJBMqConnHelper2.Init(NormalMqUrl, nil)
	if err != nil {
		logger.Errorf("Init MqConn fail,err ： %v", err)
		return
	}
	HJBMqChannel = new(mq.RabbitMqChannelHelper)
	//err = HJBMqChannel.Init("hjb", HJBMqConnHelper, registerListener, nil)
	err = HJBMqChannel.Init("hjb", HJBMqConnHelper, nil, nil)
	if err != nil {
		logger.Error("Init mq channel fail", err)
		return
	}
	HJBMqChannel2 = new(mq.RabbitMqChannelHelper)
	err = HJBMqChannel2.Init("hjb2", HJBMqConnHelper2, nil, nil)
	if err != nil {
		logger.Error("Init mq channel fail", err)
		return
	}
	DeclareHJBExchange()
	logger.Info("producer start")
	publishToExchange()
	for {
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

func publishToExchange() {
	if HJBMqChannel == nil {
		logger.Error("channel is nil")
		return
	}
	statistic.StartTime = time.Now()
	wg.Add(2)
	go run()
	go run2()
	wg.Wait()
	sendLastMsg()
	statistic.EndTime = time.Now()
	logger.Infof("总测试条数:%d,成功条数:%d ,耗时:%f /秒", *times, statistic.SuccTimes, statistic.EndTime.Sub(statistic.StartTime).Seconds())
}

func run() {
	var err error
	var msg amqp.Publishing
	var _msg _Msg
	for i := 1; i <= 50000; i++ {
		_msg.StartTime = time.Now()
		_msg.Index = i
		_msg.TotalMsg = *times
		_msg.IsEnd, _msg.IsStart = false, false
		if i == 1 {
			_msg.IsStart = true
		}
		if i == *times {
			_msg.IsEnd = true
		}
		j, _ := json.Marshal(_msg)
		msg.Body = []byte(j)
		//logger.Infof("task:%d,msg", i)
		err = HJBMqChannel.WrapPublish(HJBExchange, HJBRoutingKey, false, false, msg)
		if err != nil {
			logger.Errorf("publish fail:%d", i)
		} else {
			statistic.SuccTimes++ //多协程中这里的++是不安全的，如果上锁又会影响测试的性能。所以忽略这里的不安全，下同
		}
	}
	wg.Done()
}

func sendLastMsg() {
	var msg amqp.Publishing
	var _msg _Msg
	var err error
	_msg.StartTime = time.Now()
	_msg.Index = 100000
	_msg.TotalMsg = *times
	_msg.IsEnd, _msg.IsStart = false, false
	_msg.IsEnd = true
	j, _ := json.Marshal(_msg)
	msg.Body = []byte(j)
	err = HJBMqChannel.WrapPublish(HJBExchange, HJBRoutingKey, false, false, msg)
	if err != nil {
		logger.Errorf("publish fail:%d", _msg.Index)
	} else {
		statistic.SuccTimes++
	}
}

func run2() {
	var err error
	var msg amqp.Publishing
	var _msg _Msg
	for i := 50001; i <= *times-1; i++ {
		_msg.StartTime = time.Now()
		_msg.Index = i
		_msg.TotalMsg = *times
		_msg.IsEnd, _msg.IsStart = false, false
		if i == 1 {
			_msg.IsStart = true
		}
		if i == *times {
			_msg.IsEnd = true
		}
		j, _ := json.Marshal(_msg)
		msg.Body = []byte(j)
		//logger.Infof("task:%d,msg", i)
		err = HJBMqChannel2.WrapPublish(HJBExchange, HJBRoutingKey, false, false, msg)
		if err != nil {
			logger.Errorf("publish fail:%d", i)
		} else {
			statistic.SuccTimes++
		}
	}
	wg.Done()
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
