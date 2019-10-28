package main

import (
	"fmt"
	"time"
)

func main() {
	forever := make(chan bool)
	//utils.Run()
	start := time.Now()
	fmt.Println("start")
	ticker := time.NewTicker(time.Second * time.Duration(3))
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("done")
			return
		case t := <-ticker.C:
			fmt.Println("timeticker:", t.Sub(start).Seconds())
			return
		}
	}
	fmt.Println("stop")
	fmt.Println(time.Now().Sub(start).Seconds())
	<-forever
}
