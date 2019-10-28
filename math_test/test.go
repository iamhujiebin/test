package main

import (
	"fmt"
	"strings"
)

func main() {
	dict := get62Strs_tmp()
	origin := 53220
	var target string
	m := origin
	if origin < len(dict) {
		target = dict[origin]
		return
	}
	for m >= len(dict) {
		mod := m % len(dict)
		m = m / len(dict)
		target = fmt.Sprintf("%s%s", dict[mod], target)
	}
	if m < len(dict) {
		target = fmt.Sprintf("%s%s", dict[m], target)
	}
	fmt.Println(target)
}

func get62Strs_tmp() []string {
	str62 := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	str62Arr := make([]string, 62)
	for i := 0; i < len(str62); i++ {
		ch := str62[i]
		str62Arr[i] = strings.Replace(fmt.Sprintf("%q", ch), "'", "", -1)
	}
	for i := len(str62Arr); i > 0; i-- {
		//r := rand.Intn(i)
		//str62Arr[i-1], str62Arr[r] = str62Arr[r], str62Arr[i-1] //todo 加上这句打乱顺序
	}
	return str62Arr
}
