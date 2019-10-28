package main

import (
	"fmt"
	"./balance"
	"math/rand"
	"os"
	"time"
)

func main() {
	//初始化主机
	//insts := make([]*balance.Instance, 16) //append会自动扩容
	var insts []*balance.Instance
	for i := 0; i < 4; i++ {
		host := fmt.Sprintf("192.168.0.%d", rand.Intn(100))
		one := balance.NewInstance(host, 8080)
		insts = append(insts, one)
	}
	fmt.Printf("all insts:%v",insts)
	//var balancer balance.Balancer
	var BalancerName = "random"
	if len(os.Args) > 1 {
		BalancerName = os.Args[1]
	}
	for {
		inst, err := balance.DoBalance(BalancerName, insts)
		if err != nil {
			fmt.Println("Do balance err:", err)
			continue
		}
		fmt.Println(inst)
		time.Sleep(time.Second*3)
	}

}
