package main

import (
	"./zk_tool"
	"encoding/json"
	"fmt"
)

func main() {
	servers := []string{"192.168.16.20:2181"}
	client, err := zk_tool.NewClient(servers, "/api", 10)
	if err != nil {
		panic(err)
	}
	defer client.Close()
	nodes, err := client.GetNodes("user")
	if err != nil {
		panic(err)
	}
	for _, node := range nodes {
		j, _ := json.Marshal(node)
		fmt.Println(string(j))
	}
}
