// 发布者
package main

import (
	"encoding/json"
	"flag"
	"github.com/cihub/seelog"
	"github.com/nats-io/nats"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

var logger seelog.LoggerInterface

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

func usage() {
	log.Fatalf("Usage: nats-pub [-s server (%s)] [--tls] <subject> <msg> \n", "nats://0.0.0.0:4222")
}

func main() {
	// 下面定义连接到server的URL
	var urls = flag.String("s", "nats://192.168.16.20:4222", "The nats server URLs (separated by comma)")
	//var urls = flag.String("s", "nats://0.0.0.0:4222", "The nats server URLs (separated by comma)")
	// 是否使用TLS安全传输协议
	var tls = flag.Bool("tls", false, "Use TLS Secure Connection")
	var times = flag.Int("times", 10, "次数")

	// 下面是判断参数
	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	port := 1111
	datas, err := ioutil.ReadFile("./seelog-nats.xml")
	contentStr := string(datas)
	contentStr = strings.Replace(contentStr, "##{port}", strconv.Itoa(port), -1)
	logger, err = seelog.LoggerFromConfigAsString(contentStr)
	if err != nil {
		logger.Error(err)
		return
	}

	args := flag.Args()
	if len(args) < 1 {
		usage()
	}

	// 下面填充nats的一些选项
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
	defer nc.Close()

	// 下面定义subject和msg
	subj := args[0]
	n := 0
	for {
		n++
		log.Printf("第%d次", n)
		logger.Infof("第%d次", n)
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
			err := nc.Publish(subj, j)
			if err == nil {
				statistic.SuccTimes++
			}
		}
		time.Sleep(10 * time.Second)
	}
	statistic.EndTime = time.Now()

	// 刷新缓冲区
	nc.Flush()
	for {
	}
}
