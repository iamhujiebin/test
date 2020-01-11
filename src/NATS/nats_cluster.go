package main

import (
	"github.com/nats-io/nats"
	"nonolive/nonoutils/nonolog"
	"time"
)

func main() {
	nonolog.MainLogger().Infof("nats cluster")
	//servers := "nats://localhost:4222,nats://localhost:5222"
	servers := "nats://localhost:4222" //只配置一个，也可以动态发现到集群的其他server
	//nc, err = nats.Connect(servers, nats.DontRandomize())
	nc, err := nats.Connect(servers, nats.MaxReconnects(5), nats.ReconnectWait(2*time.Second),
		nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
			nonolog.MainLogger().Infof("Got disconnected! Reason: %q\n", err)
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			nonolog.MainLogger().Infof("Got reconnected to %v!\n", nc.ConnectedUrl())
		}),
		nats.ClosedHandler(func(nc *nats.Conn) {
			nonolog.MainLogger().Infof("Connection closed. Reason: %q\n", nc.LastError())
		}),
		nats.ErrorHandler(func(nc *nats.Conn, sub *nats.Subscription, err error) {
			if err == nats.ErrSlowConsumer {
				dropped, _ := sub.Dropped()
				nonolog.MainLogger().Errorf("Slow consumer on subject %s dropped %d messages", sub.Subject, dropped)
			}
		}),
	)
	if err != nil {
		nonolog.MainLogger().Errorf("err:%v", err)
	}
	nc.Subscribe("nats_cluster", func(msg *nats.Msg) {
		nonolog.MainLogger().Infof("data:%v", string(msg.Data))
	})
	for {
		time.Sleep(time.Second)
		nc.Publish("nats_cluster", []byte("hi"))
	}
	forever := make(chan int)
	<-forever
}
