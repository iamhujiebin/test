package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Msg struct {
	Starttime int `json:"starttime"`
	Endtime   int `json:"endtime"`
	Index     int `json:"index"`
}

func main() {
	msg := Msg{
		Starttime: time.Now().Nanosecond(),
		Endtime:   time.Now().Nanosecond() + 1111,
		Index:     111,
	}
	j, err := json.Marshal(msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(j))
	a := int(64)
	b := int(56)
	fmt.Println(float64(b) / float64(a))
}
