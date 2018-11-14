package main

import (
	dis "./discovery"
	"fmt"
	"log"
	"time"
)

func main() {
	serviceName := "s-hjb-test"
	serviceInfo := dis.ServiceInfo{
		IP:   "192.168.18.102",
		Port: 9527,
	}
	s, err := dis.NewService(serviceName, serviceInfo, []string{
		"http://192.168.16.20:2379",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("name:%s, ip:%s port:%d \n", s.Name, s.Info.IP, s.Info.Port)
	go func() {
		time.Sleep(time.Second * 60)
		s.Stop()
	}()
	s.Start()
}
