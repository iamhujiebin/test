package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var dict []string

func main() {
	urlMap := make(map[string]int)
	urlCodeMap := make(map[string]int)
	origin := 300000
	dict = get62Strs()
	for i := origin; i <= origin+1000000; i++ {
		if _, ok := urlMap[shortUrl(i)]; ok {
			fmt.Printf("caution!!---oldCode:%d,newCode:%d,url:%s\n", urlCodeMap[shortUrl(i)], i, shortUrl(i))
		}
		urlMap[shortUrl(i)]++
		urlCodeMap[shortUrl(i)] = i
		fmt.Println(i, shortUrl(i))
	}
	for k, v := range urlMap {
		if v > 1 {
			fmt.Println(k, v, urlCodeMap[k])
		}
	}
	fmt.Println("finish")
}

func shortUrl(origin int) (target string) {
	//此算法参考decimal.go这个文件
	m := origin
	num := 0
	x := 1
	for m/10 != 0 {
		num = num*10 + m%10
		m = m / 10
		x = x * 10
	}
	num = m*x + num
	if origin < len(dict) {
		target = dict[origin]
		return
	}
	for num >= len(dict) {
		mod := num % len(dict)
		num = num / len(dict)
		target = fmt.Sprintf("%s%s", dict[mod], target)
	}
	if num < len(dict) {
		target = fmt.Sprintf("%s%s", dict[num], target)
	}
	return
}

func get62Strs() []string {
	str62 := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	str62Arr := make([]string, 62)
	for i := 0; i < len(str62); i++ {
		ch := str62[i]
		str62Arr[i] = strings.Replace(fmt.Sprintf("%q", ch), "'", "", -1)
	}
	for i := len(str62Arr); i > 0; i-- {
		r := rand.Intn(i)
		str62Arr[i-1], str62Arr[r] = str62Arr[r], str62Arr[i-1] //加上这句打乱顺序
	}
	return str62Arr
}
