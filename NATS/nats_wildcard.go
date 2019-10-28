package main

import (
	"github.com/nats-io/nats"
	"nonolive/nonoutils/nonolog"
)

func main() {
	nonolog.MainLogger().Infof("nats wildcard")
	nc, _ := nats.Connect("127.0.0.1:4222")

	nc.Subscribe("jie.my.test", func(msg *nats.Msg) {
		nonolog.MainLogger().Infof("jie.my.test match")
	})
	nc.Subscribe("jie.*.test", func(msg *nats.Msg) {
		nonolog.MainLogger().Infof("jie.*.test match")
	})
	nc.Subscribe("jie.*", func(msg *nats.Msg) {
		nonolog.MainLogger().Infof("jie* match")
	})
	nc.Subscribe("jie.>", func(msg *nats.Msg) {
		nonolog.MainLogger().Infof("jie.> match")
	})
	nc.Subscribe("jie>", func(msg *nats.Msg) {
		nonolog.MainLogger().Infof("jie> match :")
	})
	nc.Subscribe("jie.*.tes", func(msg *nats.Msg) {
		nonolog.MainLogger().Infof("jie.*.tes match :")
	})
	//nc.Subscribe("jie.>.test", func(msg *nats.Msg) { //非法的通配符。导致所有都收不到？
	//	nonolog.MainLogger().Infof("jie.>.test match :")
	//})

	nc.Publish("jie.my.test", []byte("hi"))
	forever := make(chan int)
	<-forever
}
