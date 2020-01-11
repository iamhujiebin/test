package main

import "fmt"

type MyStruct struct {
	i int
}

func myFunction(b *MyStruct) {
	//b.i = 41
	fmt.Printf("in my_function - b=(%v, %p)\n", b, &b)
}

func main() {
	b := &MyStruct{i: 40}
	fmt.Printf("before calling - b=(%v, %p)\n", b, &b)
	myFunction(b)
	fmt.Printf("after calling  -  b=(%v, %p)\n", b, &b)
}
