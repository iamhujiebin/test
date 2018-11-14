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
		//printMsg(msg, i)
	})

	// 下面设置自动"反注册", 当达到一定的数量后就自动执行"发注册"
	//const MAX_WANTED = 3
	//sub.AutoUnsubscribe(MAX_WANTED)
	go echoTotal()

	log.Printf("Listening on [%s]\n", subj)
	runtime.Goexit()
}

func echoTotal() {
	for {
		time.Sleep(5 * time.Second)
		log.Printf("total:%d\n", total)
	}
}
