// 订阅者
package main

import (
	"encoding/json"
	"flag"
	"github.com/nats-io/nats"
	"log"
	"runtime"
	"strings"
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

var total int
var statistic _Statistic

func usage() {
	log.Fatalf("Usage: nats-sub [-s server] [--tls] [-t] <subject> \n")
}

func printMsg(m *nats.Msg, i int) {
	log.Printf("[#%d] Received on [%s]: '%s'\n", i, m.Subject, string(m.Data))
}

func main() {
	var urls = flag.String("s", "nats://192.168.16.20:4222", "The nats server URLs (separated by comma)")
	var tls = flag.Bool("tls", false, "Use Secure Connection")

	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		usage()
	}

	// 解析可选项
	opts := nats.DefaultOptions
	opts.Servers = strings.Split(*urls, ",")
	for i, s := range opts.Servers {
		opts.Servers[i] = strings.Trim(s, " ")
	}
	opts.Secure = *tls

	// 连接到gnatsd
	nc, err := opts.Connect()
	if err != nil {
		log.Fatalf("Can't connect: %v\n", err)
	}

	// 订阅的subject
	subj, i := args[0], 0

	// 订阅主题, 当收到subject时候执行后面的func函数
	// 返回值sub是subscription的实例
	//
	//sub, _ := nc.Subscribe(subj, func(msg *nats.Msg) {
	nc.Subscribe(subj, func(msg *nats.Msg) {
		i += 1
		total += 1
		statistic.SuccTimes++
		var _msg _Msg
		json.Unmarshal(msg.Data, &_msg)
		_msg.EndTime = time.Now()
		statistic.SuccMsgDuration += _msg.EndTime.Sub(_msg.StartTime).Seconds()
		statistic.TotalMsg = _msg.TotalMsg
		if _msg.IsStart {
			statistic.StartTime = time.Now()
			log.Printf("startime:%d", _msg.StartTime.Unix())
		}
		if _msg.IsEnd {
			statistic.EndTime = time.Now()
			statistic.IsEnd = true
			log.Printf("endtime:%d", _msg.EndTime.Unix())
		}
		//log.Println(_msg.Index)
		//printMsg(msg, i)
	})

	// 下面设置取消订阅 当达到一定的数量后就自动执行取消订阅
	//const MAX_WANTED = 3
	//sub.AutoUnsubscribe(MAX_WANTED)
	go echoTotal()

	log.Printf("Listening on [%s]\n", subj)
	runtime.Goexit()
}

func echoTotal() {
	for {
		time.Sleep(5 * time.Second)
		if statistic.IsEnd {
			log.Printf("subscribe:总测试条数:%d,成功条数:%d ,成功率:%0.3f,总耗时:%f /秒,一条消息耗时:%f/秒 \n",
				statistic.TotalMsg, statistic.SuccTimes, float64(statistic.SuccTimes)/float64(statistic.TotalMsg),
				statistic.EndTime.Sub(statistic.StartTime).Seconds(), statistic.SuccMsgDuration/float64(statistic.SuccTimes))
			statistic = _Statistic{}
		}
	}
}
