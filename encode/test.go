package main

import (
	"fmt"
	"net/url"
)

func main() {
	t, _ := url.ParseQuery("%EF%B8%8F%E2%80%8D%EF%B8%8F%E2%80%8D%EF%B8%8F%E2%80%8D")
	fmt.Println()
}
