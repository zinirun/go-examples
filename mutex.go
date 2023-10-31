package main

import (
	"fmt"
	"sync"
)

var count int
var mu sync.Mutex

func increment() {
	mu.Lock()
	count++
	mu.Unlock()
}

func Mutex() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	wg.Wait()
	fmt.Println("Count:", count)
}
