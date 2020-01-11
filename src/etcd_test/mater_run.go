package main

import (
	dis "./discovery"
	"fmt"
	"log"
	"time"
)

func main() {
	m, err := dis.NewMaster([]string{
		"http://192.168.16.20:2379",
	}, "services/")
	if err != nil {
		log.Fatal(err)
	}
	for {
		for k, v := range m.Nodes {
			fmt.Printf("nodes:%s,ip=%s,port=%d", k, v.Info.IP, v.Info.Port)
		}
		fmt.Printf("nodes num = %d\n", len(m.Nodes))
		time.Sleep(time.Second * 5)
	}
}
