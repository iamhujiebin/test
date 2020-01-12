package main

import "fmt"

func main() {
	s := "11弟弟"
	t := []rune(s)
	var str string
	for i, v := range t {
		str = str + string(v)
		fmt.Println(i, str)
	}
}
