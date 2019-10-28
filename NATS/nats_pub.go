package main

import (
	"github.com/nats-io/nats"
	"nonolive/nonoutils/nonolog"
	"time"
)

func main() {
	nc, _ := nats.Connect("127.0.0.1:4222")
	//nc, _ := nats.Connect("192.168.16.27:4222")
	//simple test
	err := nc.Publish("test1", []byte("hello world!"))
	if err != nil {
		nonolog.MainLogger().Errorf("err:%s", err.Error())
	}
	//reply test
	msg, err := nc.Request("help", []byte("help me"), 10*time.Second)
	if err != nil {
		nonolog.MainLogger().Errorf("err:%s", err.Error())
	} else {
		nonolog.MainLogger().Errorf("msg:%s", msg.Data)
	}
	//queue test
	nc.Publish("queue_test", []byte("hello queue!"))
	if err != nil {
		nonolog.MainLogger().Errorf("err:%s", err.Error())
	}
	forever := make(chan int)
	<-forever
	defer nc.Close()
}
