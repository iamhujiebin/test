package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
	"net"
	"time"
)

const RECV_BUF_LEN = 1024

func main() {
	client, err := api.NewClient(api.DefaultConfig())

	if err != nil {
		log.Fatal("consul client error :", err)
	}

	for {
		time.Sleep(time.Duration(int64(3)) * time.Second)
		var services map[string]*api.AgentService
		var err error
		services, err = client.Agent().Services()
		if err != nil {
			log.Println("in consul list services :", err)
		}
		if _, found := services["serverNode_1"]; !found {
			log.Println("serverNode_1 not found")
			continue
		}
		sendData(services["serverNode_1"])
	}
}

func sendData(service *api.AgentService) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", service.Address, service.Port))
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	buf := make([]byte, RECV_BUF_LEN)
	i := 0
	for {
		i++
		msg := fmt.Sprintf("hello world,%03d", i)
		n, err := conn.Write([]byte(msg))
		if err != nil {
			println("Write Buffer Error:", err.Error())
			break
		}
		n, err = conn.Read(buf)
		if err != nil {
			println("Read Buffer Error:", err.Error())
			break
		}
		log.Println("get:", string(buf[0:n]))

		//等一秒钟
		time.Sleep(time.Second)
	}
}
