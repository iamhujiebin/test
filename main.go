package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type A struct {
		Code int  `json:"code,string"`
		Name *int `json:"name,string"`
	}
	str := `{"code":"101","name":"99"}`
	str = `{"code":"101"}`
	a := new(A)
	json.Unmarshal([]byte(str), &a)
	fmt.Println(a.Code, a.Name)
}
