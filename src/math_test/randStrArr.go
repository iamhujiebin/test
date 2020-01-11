package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	str62 := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	fmt.Println(str62)
	str62Arr := make([]string, 62)
	for i := 0; i < len(str62); i++ {
		ch := str62[i]
		str62Arr[i] = strings.Replace(fmt.Sprintf("%q", ch), "'", "", -1)
	}
	fmt.Println(str62Arr)
	fmt.Println(len(str62Arr))
	for i := len(str62Arr); i > 0; i-- {
		r := rand.Intn(i)
		str62Arr[i-1], str62Arr[r] = str62Arr[r], str62Arr[i-1]
	}
	fmt.Println(str62Arr)
	str62New := strings.Join(str62Arr, "")
	fmt.Println(str62New)
}
