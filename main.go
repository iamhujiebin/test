package main

import (
	"fmt"
	"time"
)

func main() {
	date := time.Date(2020, 1, 31, 0, 0, 0, 0, time.UTC)
	fmt.Println(date)
	b := date.AddDate(0, 2, 0)
	fmt.Println(b)
	fmt.Println(b.Sub(date).Hours())
}
