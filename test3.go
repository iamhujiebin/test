package main

import (
	"flamingo/mytest"
	"fmt"
)

func main() {
	fmt.Println("main started")
	f()
	fmt.Println("Returned normally from f.")
}

func init() {
	fmt.Println("iam there init func")
}

func f() {
	mytest.MyTest()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
		fmt.Println("just little defer func")
	}()
	//panic("hehe")
	fmt.Println("Calling g.")
	g()
	fmt.Println("Returned normally from g.")
}

func g() {
	//panic("ERROR")
	fmt.Println("no panic")
}
