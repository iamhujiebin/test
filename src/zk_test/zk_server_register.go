package main

import (
	"./zk_tool"
)

func main() {
	servers := []string{"192.168.16.20:2181"}
	client, err := zk_tool.NewClient(servers, "/api", 10)
	if err != nil {
		panic(err)
	}
	defer client.Close()
	node1 := &zk_tool.ServiceNode{"user", "127.0.0.1", 4000}
	node2 := &zk_tool.ServiceNode{"user", "127.0.0.1", 4001}
	node3 := &zk_tool.ServiceNode{"user", "127.0.0.1", 4002}
	if err := client.Register(node1); err != nil {
		panic(err)
	}
	if err := client.Register(node2); err != nil {
		panic(err)
	}
	if err := client.Register(node3); err != nil {
		panic(err)
	}
	for {

	}
}
