package main

import (
	"encoding/json"
	"flag"
	"github.com/nats-io/go-nats-streaming"
	"log"
	"runtime"
	"time"
)

type _Statistic struct {
	StartTime time.Time
	EndTime   time.Time
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

var statistic _Statistic

func main() {
	var times = flag.Int("times", 10, "次数")
	flag.Parse()
	//stan.Connect(clusterID, clientID, ops ...Option)
	//預設clusterID為test-cluster
	ns, err := stan.Connect("test-cluster", "myid", stan.NatsURL("nats://0.0.0.0:4222"))
	if err != nil {
		panic(err)
	}
	// Simple Synchronous Publisher
	// does not return until an ack has been received from NATS Streaming
	//釋出50000條訊息
	statistic.StartTime = time.Now()
	for i := 1; i <= *times; i++ {
		msg := _Msg{
			Index:     i,
			StartTime: time.Now(),
			TotalMsg:  *times,
		}
		if i == 1 {
			msg.IsStart = true
		}
		if i == *times {
			msg.IsEnd = true
		}
		j, _ := json.Marshal(msg)
		err := ns.Publish("logp", j)
		if err == nil {
			statistic.SuccTimes++
		}
	}
	statistic.EndTime = time.Now()
	log.Printf("publish:总测试条数:%d,成功条数:%d ,耗时:%f /秒\n", *times, statistic.SuccTimes, statistic.EndTime.Sub(statistic.StartTime).Seconds())
	ns.Close()
	runtime.Goexit()
}
