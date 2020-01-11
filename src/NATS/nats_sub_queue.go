package main

import (
	"github.com/nats-io/nats"
	"nonolive/nonoutils/nonolog"
)

func main() {
	nonolog.MainLogger().Infof("start nats sub queue")
	nc, _ := nats.Connect("127.0.0.1:4222")
	nc.QueueSubscribe("queue_test", "queue", func(msg *nats.Msg) {
		nonolog.MainLogger().Infof("queue_1 receive:%s", string(msg.Data))
	})
	nc.QueueSubscribe("queue_test", "queue", func(msg *nats.Msg) {
		nonolog.MainLogger().Infof("queue_2 receive:%s", string(msg.Data))
	})
	ch := make(chan *nats.Msg, 64)
	nc.ChanQueueSubscribe("queue_test", "queue", ch)
	go chanQueueMsg(ch)
	nc.QueueSubscribe("queue_test", "other_queue", func(msg *nats.Msg) {
		nonolog.MainLogger().Infof("other queue receive:%s", string(msg.Data))
	})
	forever := make(chan int)
	<-forever
}

func chanQueueMsg(ch chan *nats.Msg) {
	for {
		select {
		case m := <-ch:
			nonolog.MainLogger().Infof("queue_chan receive:%s", string(m.Data))
		}
	}
	close(ch)
}
