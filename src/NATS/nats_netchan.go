package main

import (
	"github.com/nats-io/nats"
	"nonolive/nonoutils/nonolog"
)

type User struct {
	UserId int
	Name   string
}

func main() {

	nonolog.MainLogger().Infof("start nats netChan")
	nc, _ := nats.Connect("127.0.0.1:4222")
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)

	sendCh, receiveCh := make(chan *User), make(chan *User)
	c.BindSendChan("netChan", sendCh)
	c.BindRecvChan("netChan", receiveCh)
	go netChan(receiveCh)
	sendCh <- &User{UserId: 2, Name: "my"}
	forever := make(chan int)
	<-forever
}

func netChan(ch chan *User) {
	for {
		select {
		case who := <-ch:
			nonolog.MainLogger().Infof("who:%v", who)
		}
	}
}
