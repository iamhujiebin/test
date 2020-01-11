package main

import (
	"encoding/json"
	"flag"
	"github.com/nats-io/nats"
	"log"
	"nonolive/nonoutils/nononats"
	//"runtime"
	"time"
)

type _Statistic struct {
	Starttime time.Time
	Endtime   time.Time
	TotalMsg  int
	ReplyMsg  int
}

type _Msg struct {
	IsStart  bool
	IsEnd    bool
	Index    int
	TotalMsg int
}

var statistic _Statistic

func main() {
	var forever chan int
	times := flag.Int("times", 1, "测试数量")
	flag.Parse()
	helper := new(nononats.NatsConnHelper)
	urls := []string{"nats://192.168.16.20:4222"}
	helper.SetIsMonitor(false)
	helper.Init(urls, nil, nil)
	helper.SetReplyCallBack("reply", replyCallback)
	statistic.TotalMsg = *times
	n := 0
	for {
		n++
		log.Printf("第%d次测试", n)
		for i := 1; i <= *times; i++ {
			_msg := new(_Msg)
			_msg.Index = n
			_msg.TotalMsg = *times
			if i == 1 {
				_msg.IsStart = true
			}
			if i == *times {
				//_msg.IsEnd = true
			}
			j, _ := json.Marshal(_msg)
			_, err := helper.WrapPublish("hjb", "reply", j)
			if err != nil {
				log.Println(err)
			}
		}
		time.Sleep(time.Second)
	}
	//runtime.Goexit()
	<-forever
	helper.Close()
}

func replyCallback(msg *nats.Msg) {
	//log.Printf("callback come reply:%s", string(msg.Data))
	statistic.ReplyMsg++
	var nonoReplyMsg nononats.ReplyMsg
	var _msg _Msg
	json.Unmarshal(msg.Data, &nonoReplyMsg)
	json.Unmarshal(nonoReplyMsg.Nmsg.Data, &_msg)
	if _msg.IsStart {
		statistic.Starttime = time.Now()
	}
	log.Printf("callback come reply:%d", _msg.Index)
	if _msg.IsEnd {
		statistic.Endtime = time.Now()
		log.Printf("总测试次数:%d,成功reply次数:%d, 耗时:%f /秒\n", statistic.TotalMsg, statistic.ReplyMsg,
			statistic.Endtime.Sub(statistic.Starttime).Seconds())
		statistic = _Statistic{}
	}
}
