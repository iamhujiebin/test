package main

import (
	"encoding/json"
	"github.com/nats-io/nats"
	"nonolive/nonoutils/nonolog"
	"time"
)

func main() {
	nonolog.MainLogger().Infof("start nats sub reply")
	nc, _ := nats.Connect("127.0.0.1:4222")
	sub, err := nc.SubscribeSync("help") //阻塞监听者(同步)
	m, err := sub.NextMsg(time.Second * 10)
	if err != nil {
		nonolog.MainLogger().Errorf("error:%s", err.Error())
	} else {
		nc.Publish(m.Reply, []byte("i can't help sync"))
		nonolog.MainLogger().Infof("sync sub msg:%s\ndata:%v", m, string(m.Data))
	}
	sub.Unsubscribe()
	nonolog.MainLogger().Infof("after sync")
	/*---------------------chan 订阅---------------------*/
	ch := make(chan *nats.Msg, 64)
	sub, err = nc.ChanSubscribe("help", ch)
	go chanMsg(sub, ch) //注释掉这行，可以看到两个普通的订阅者，reply速度随机
	/*---------------------chan 订阅---------------------*/

	/*---------------------异步非阻塞监听者----------------*/
	nc.Subscribe("help", func(msg *nats.Msg) {
		msg.Respond([]byte("i can't help"))
		j, _ := json.Marshal(msg)
		nonolog.MainLogger().Infof("receive1:%v\ndata:%v", string(j), string(msg.Data))
	})
	nc.Subscribe("help", func(msg *nats.Msg) {
		msg.Respond([]byte("i can't help2"))
		j, _ := json.Marshal(msg)
		nonolog.MainLogger().Infof("receive2:%v\ndata:%v", string(j), string(msg.Data))
	})
	/*---------------------异步非阻塞监听者----------------*/

	forever := make(chan int)
	<-forever
}

func chanMsg(sub *nats.Subscription, ch chan *nats.Msg) {
EXIT:
	for {
		select {
		case m := <-ch:
			nonolog.MainLogger().Infof("chan sub msg:%s\ndata:%v", m, string(m.Data))
			m.Respond([]byte("i can't help chan"))
			sub.Drain()       // 处理完callback,但是不会接受新的消息
			sub.Unsubscribe() //不会接受新的消息
			break EXIT
		}
	}
	//取消订阅之后，应该close掉chan。
	close(ch)
}
