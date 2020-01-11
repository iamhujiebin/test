// 订阅者
package main

import (
	"flag"
	"github.com/nats-io/nats"
	"log"
	"runtime"
	"strings"
	"time"
)

var total int
var total2 int

func usage() {
	log.Fatalf("Usage: nats-sub [-s server] [--tls] [-t] <subject> \n")
}

func printMsg(m *nats.Msg, i int) {
	log.Printf("[#consumer:%d] Received on [%s]: '%s'\n", i, m.Subject, string(m.Data))
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
	subj := args[0]
	log.Printf("Listening on [%s]\n", subj)
	//go echoTotal()

	/*
		//sub, _ := nc.Subscribe(subj, func(msg *nats.Msg) {
		nc.Subscribe(subj, func(msg *nats.Msg) {
				total += 1
		})
	*/
	/*
		// 下面设置自动取消订阅，如果使用ChanSubscribe，就需要ch.Close()
		const MAX_WANTED = 3
		sub.AutoUnsubscribe(MAX_WANTED)
	*/
	/*
		ch := make(chan *nats.Msg, 1000) //使用chan的话，要注意缓冲区的问题。缓冲区小，并发高，会导致丢包？
		nc.ChanSubscribe(subj, ch)
		for range ch {
			total2++
		}
	*/
	/*
		nc.QueueSubscribe(subj, "hjb_queue", func(msg *nats.Msg) {
			total++
			//printMsg(msg, 1)
		})
		nc.QueueSubscribe(subj, "hjb_queue", func(msg *nats.Msg) {
			total2++
			//printMsg(msg, 2)
		})
	*/
	/*
		ch1, ch2 := make(chan *nats.Msg, 1000), make(chan *nats.Msg, 1000)
		nc.ChanQueueSubscribe(subj, "hjb_queue", ch1)
		nc.ChanQueueSubscribe(subj, "hjb_queue2", ch2)
		for {
			select {
			case <-ch1:
				total++
			case <-ch2:
				total2++
			}
		}
	*/
	/*
		sub, _ := nc.QueueSubscribeSync(subj, "hjb_queue")
		for msg, _ := sub.NextMsg(time.Second * 100); msg != nil; msg, _ = sub.NextMsg(time.Second * 100) {
			printMsg(msg, 0)
		}
	*/
	/*
		ch := make(chan *nats.Msg, 1000)
		nc.QueueSubscribeSyncWithChan(subj, "hjb_queue", ch)
		for range ch {
			total++
		}
	*/
	nc.Subscribe(subj, func(m *nats.Msg) {
		log.Printf("d:%s\n r:%s", string(m.Data), m.Reply)
		nc.Publish(m.Reply, []byte("acks"))
	})

	runtime.Goexit()
}

func echoTotal() {
	for {
		time.Sleep(1 * time.Second)
		log.Printf("total:%d -- total2:%d\n", total, total2)
	}
}
