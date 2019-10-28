package services

import (
	"fmt"
)

type Service struct {
	Name string
	Id   int
}

func ServiceFunc() {
	sname := "servie name"
	fmt.Printf("service func :%s\n", sname)
}
