package main

import (
	"encoding/json"
	"github.com/nats-io/nats"
	"nonolive/nonoutils/nonolog"
)

func main() {
	nonolog.MainLogger().Infof("start nats sub")
	nc, _ := nats.Connect("127.0.0.1:4222")
	//nc, _ := nats.Connect("192.168.16.27:4222")
	nc.Subscribe("test1", func(msg *nats.Msg) {
		j, _ := json.Marshal(msg)
		nonolog.MainLogger().Infof("receive:%v\ndata:%v", string(j), string(msg.Data))
	})
	forever := make(chan int)
	<-forever
}
