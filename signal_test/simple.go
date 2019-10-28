package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

// 监听指定信号
func main() {

	fmt.Println("启动")
	go ExceptionExit()
	runtime.Goexit()
}

func ExceptionExit() {
	//合建chan
	c := make(chan os.Signal)
	//监听指定信号 ctrl+c kill
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	//阻塞直至有信号传入
	s := <-c
	fmt.Println("退出信号", s)
	os.Exit(0)
}
