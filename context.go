package main

import (
	"context"
	"fmt"
	"time"
)

func Context() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(2 * time.Second)
		cancel() // After 2 seconds, cancel the context
	}()

	select {
	case <-ctx.Done():
		fmt.Println("Context was canceled")
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout")
	}
}
