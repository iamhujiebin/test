package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"io"
	"log"
	"net"
	"net/http"
)

const RECV_BUF_LEN = 1024

func consulCheck(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.URL.Path)
	fmt.Fprintln(w, "consultCheck")
}

func registerServer() {
	config := api.DefaultConfig()
	//config.Address = "192.168.16.20:8500"
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatal("consul client error:", err)
	}
	checkPort := 9000
	registration := new(api.AgentServiceRegistration)
	registration.ID = "serverNode_1"
	registration.Name = "serverNode"
	registration.Port = 9527
	registration.Tags = []string{"serverNode"}
	//registration.Address = "127.0.0.1"
	registration.Address = "192.168.18.102"
	registration.Check = &api.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%s:%d/%s", registration.Address, checkPort, "check"),
		Timeout:                        "3s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "30s",
	}
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatal("register server error:", err)
	}

	http.HandleFunc("/check", consulCheck)
	http.ListenAndServe(fmt.Sprintf(":%d", checkPort), nil)
}

func main() {
	go registerServer()

	ln, err := net.Listen("tcp", "0.0.0.0:9527")

	if err != nil {
		panic("Error:" + err.Error())
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic("Error:" + err.Error())
		}
		go EchoServer(conn)
	}
}

func EchoServer(conn net.Conn) {
	buf := make([]byte, RECV_BUF_LEN)
	defer conn.Close()

	for {
		n, err := conn.Read(buf)
		switch err {
		case nil:
			log.Println("get and echo:", "EchoServer"+string(buf[0:n]))
			conn.Write(append([]byte("EchoServer"), buf[0:n]...))
		case io.EOF:
			log.Println("Warning: End of data: %s\n", err)
			return
		default:
			log.Printf("Error: Reading data: %s\n", err)
			return
		}
	}
}
