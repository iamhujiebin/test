package main

import "fmt"

func main() {
	origin := 9999
	var target string
	defer func() {
		fmt.Printf("origin:%d\ntarget:%s", origin, target)
	}()
	dict := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
	m := origin
	if origin < len(dict) {
		target = dict[origin]
		return
	}
	for m > len(dict) {
		mod := m % len(dict)
		m = m / len(dict)
		target = fmt.Sprintf("%s%s", dict[mod], target)
	}
	if m < len(dict) {
		target = fmt.Sprintf("%s%s", dict[m], target)
	}
}
