package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	//str := "2018-12-11T17:36:58.0328758+08:00"
	//t2, _ := time.Parse("2018-12-11T17:36:58.0328758+08:00", str)
	//fmt.Println(t2)
	//j, _ := json.Marshal(time.Now())
	//fmt.Println(string(j))
	var t3 time.Time
	j := []byte("2018-12-11T20:16:47.3140431+08:00")
	json.Unmarshal(j, &t3)
	fmt.Println(t3)
}
