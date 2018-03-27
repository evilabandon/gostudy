package main

import (
	"context"
	"fmt"
	"time"
)

func run(ctx context.Context, threadId int) {
	for {
		select {
		case <- ctx.Done():
			fmt.Println("timeout")
		default:
			fmt.Println("thred running ", threadId)
		}

	}
}

func timeout(cancel context.CancelFunc) {
	time.Sleep(1 * time.Second)
	cancel()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go run(ctx, 1)
	go run(ctx, 2)
	go run(ctx, 3)
	go timeout(cancel)
	time.Sleep(2 * time.Second)
}

