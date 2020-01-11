package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		time.Sleep(time.Second)
		//ch <- 1
		//close(ch)
	}()
	fmt.Println("wait")
	<-ch
	fmt.Println("end")
}
