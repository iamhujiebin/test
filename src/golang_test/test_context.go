package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go handle(ctx, 500*time.Millisecond)

	select {
	case <-ctx.Done():
		fmt.Println("main222", ctx.Err())
	}
	time.Sleep(time.Second * 5)
}

func handle(ctx context.Context, duration time.Duration) {
loop:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("handle", ctx.Err())
			break loop
		case <-time.After(duration):
			fmt.Println("process request with", duration)
		}
	}
}
