package main

import (
	"encoding/json"
	"github.com/nats-io/nats"
	"log"
	"nonolive/nonoutils/nononats"
	"runtime"
	"time"
)

type _Statistic struct {
	Starttime  time.Time
	Endtime    time.Time
	TotalMsg   int
	RecieveMsg int
}

type _Msg struct {
	IsStart  bool
	IsEnd    bool
	Index    int
	TotalMsg int
}

var statistic _Statistic

func main() {
	helper := new(nononats.NatsConnHelper)
	urls := []string{"nats://192.168.16.20:4222"}
	helper.Init(urls, nil, nil)
	helper.SetMonitorCB(monitorCallback)
	helper.SetReconnectHandler(reconnnectCallback)
	replyMsg := &nononats.ReplyMsg{
		Msg: []byte("ack"),
	}
	subj, queue := "hjb", "queue"
	sub, _ := helper.WrapConsume(subj, queue, replyMsg, consumeCallback)
	log.Printf("subinfo:%v", sub)
	runtime.Goexit()
	helper.Close()
}

func consumeCallback(msg *nats.Msg) {
	log.Printf("iam consumer recieve msg:%s\n", string(msg.Data))
	var _msg _Msg
	json.Unmarshal(msg.Data, &_msg)
	statistic.TotalMsg = _msg.TotalMsg
	statistic.RecieveMsg++
	if _msg.IsStart {
		statistic.Starttime = time.Now()
	}
	if _msg.IsEnd {
		statistic.Endtime = time.Now()
		log.Printf("总测试次数:%d,成功consume次数:%d, 耗时:%f /秒\n", statistic.TotalMsg, statistic.RecieveMsg,
			statistic.Endtime.Sub(statistic.Starttime).Seconds())
		statistic = _Statistic{}
	}
}

func monitorCallback(conn *nats.Conn) {
	log.Printf("monitorCB---conn status:%d", conn.Status())
}

func reconnnectCallback(conn *nats.Conn) {
	log.Printf("reconnnectCallback---conn status:%d", conn.Status())
}
