package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2/bson"
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

	bj := bson.NewObjectId()
	h := bj.Hex()
	fmt.Println(h)
	fmt.Println(bson.ObjectIdHex("5bf4c2b160095251b00c7d05"))
}
