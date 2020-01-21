package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Unix(int64(1579004840), 0)
	t.In(time.UTC)
	fmt.Println(t.In(time.UTC).Location(), t.In(time.UTC).Hour())
}
