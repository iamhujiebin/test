package main

import (
	"encoding/json"
	"github.com/nats-io/go-nats-streaming"
	"log"
	"runtime"
	"time"
)

type _Statistic struct {
	StartTime       time.Time
	EndTime         time.Time
	SuccTimes       int
	IsEnd           bool
	TotalMsg        int
	SuccMsgDuration float64
}

type _Msg struct {
	StartTime time.Time `json:"starttime"`
	EndTime   time.Time `json:"endtime"`
	Index     int       `json:"index"`
	TotalMsg  int       `json:"total_msg"`
	IsStart   bool      `json:"is_start"`
	IsEnd     bool      `json:"is_end"`
}

var statistic _Statistic

func main() {
	//stan.Connect(clusterID, clientID, ops ...Option)
	ns, err := stan.Connect("test-cluster", "myid1", stan.NatsURL("nats://0.0.0.0:4222"))
	if err != nil {
		panic(err)
	}
	// Simple Synchronous Publisher
	// does not return until an ack has been received from NATS Streaming
	_, err1 := ns.Subscribe("logp", func(msg *stan.Msg) {
		statistic.SuccTimes++
		var _msg _Msg
		json.Unmarshal(msg.Data, &_msg)
		statistic.TotalMsg = _msg.TotalMsg
		_msg.EndTime = time.Now()
		statistic.SuccMsgDuration += _msg.EndTime.Sub(_msg.StartTime).Seconds()
		if _msg.IsStart {
			statistic.StartTime = time.Now()
			log.Printf("startime:%d", _msg.StartTime.Unix())
		}
		if _msg.IsEnd {
			statistic.EndTime = time.Now()
			statistic.IsEnd = true
			log.Printf("endtime:%d", _msg.EndTime.Unix())
		}
	}, stan.DurableName("cdn1"))
	if err1 != nil {
		panic(err1)
	}
	log.Printf("Listening on [%s]\n", "logp")
	go echoTotal()
	runtime.Goexit()
}

func echoTotal() {
	for {
		time.Sleep(5 * time.Second)
		if statistic.IsEnd {
			log.Printf("subscribe:总测试条数:%d,成功条数:%d ,成功率:%0.3f,总耗时:%f /秒,一条消息耗时:%f /秒 \n",
				statistic.TotalMsg, statistic.SuccTimes, float64(statistic.SuccTimes)/float64(statistic.TotalMsg),
				statistic.EndTime.Sub(statistic.StartTime).Seconds(), statistic.SuccMsgDuration/float64(statistic.SuccTimes))
			statistic = _Statistic{}
		}
	}
}
