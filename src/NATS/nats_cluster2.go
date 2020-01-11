package main

import (
	"github.com/nats-io/nats"
	"nonolive/nonoutils/nonolog"
	"time"
)

func main() {
	opts := &nats.DefaultOptions
	opts.Servers = []string{"nats://localhost:4222"}
	nc, _ := opts.Connect()
	nc.SetClosedHandler(func(nc *nats.Conn) {
		nonolog.MainLogger().Infof("Connection closed. Reason: %q\n", nc.LastError())
	})
	nc.SetDisconnectErrHandler(func(nc *nats.Conn, err error) {
		nonolog.MainLogger().Infof("Got disconnected! Reason: %q\n", err)
	})
	nc.SetReconnectHandler(func(nc *nats.Conn) {
		nonolog.MainLogger().Infof("Got reconnected to %v!\n", nc.ConnectedUrl())
	})
	nc.Subscribe("nats_cluster", func(msg *nats.Msg) {
		nonolog.MainLogger().Infof("data:%v", string(msg.Data))
	})
	for {
		time.Sleep(time.Second)
		nc.Publish("nats_cluster", []byte("hi"))
	}
}
