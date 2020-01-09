package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "jiejie110"
	fmt.Println(strings.Contains(str, "ji110"))
	fmt.Println(strings.Contains(str, "jiejie"))
}
