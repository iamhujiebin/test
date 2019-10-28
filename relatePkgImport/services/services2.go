package services

import (
	"fmt"
)

func ServiceFunc2() {
	ServiceFunc()
	s := Service{
		Name: "name",
		Id:   1,
	}
	fmt.Println(s)
}
