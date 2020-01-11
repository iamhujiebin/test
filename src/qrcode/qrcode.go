package main

import (
	"github.com/skip2/go-qrcode"
)

func main() {
	qrcode.WriteFile("http://192.168.16.20:15672/", qrcode.Highest, 512, "./mq.png")
	q := qrcode.New(content, level)
}
