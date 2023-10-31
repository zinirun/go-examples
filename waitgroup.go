package main

import (
	"fmt"
	"sync"
)

func workerWG(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	// Do work...
	fmt.Printf("Worker %d done\n", id)
}

func WaitGroup() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go workerWG(i, &wg)
	}

	wg.Wait()
	fmt.Println("All workers finished.")
}
