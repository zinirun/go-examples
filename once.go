package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func initialize() {
	fmt.Println("Initialization complete.")
}

func workerOnce(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	once.Do(initialize)
	fmt.Printf("Worker %d done\n", id)
}

func Once() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go workerOnce(i, &wg)
	}

	wg.Wait()
	fmt.Println("All workers finished.")
}
