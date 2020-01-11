package main

import (
	"github.com/nats-io/nats"
	"nonolive/nonoutils/nonolog"
)

func main() {
	type User struct {
		UserId int
		Name   string
	}
	nonolog.MainLogger().Infof("start nats encoding conn")
	nc, _ := nats.Connect("127.0.0.1:4222")
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()

	me := User{
		UserId: 1,
		Name:   "jiebin",
	}
	//subscribe
	c.Subscribe("encode_json", func(p *User) {
		nonolog.MainLogger().Infof("receive:%v", p)
	})
	//publish
	c.Publish("encode_json", &me)
	forever := make(chan int)
	<-forever

}
