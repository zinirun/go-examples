package main

import (
	"fmt"
	"sync"
	"time"
)

type Item struct {
	inStock bool
}

func (item *Item) buy() {
	if item.inStock {
		fmt.Println("Purchased item")
		item.inStock = false
	} else {
		fmt.Println("No stock")
	}
}

func Cond() {
	// 아이템 초기화
	item := &Item{}

	// 조건 변수 초기화
	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	// 고루틴이 조건을 기다리도록 함
	go func() {
		mu.Lock()
		for !item.inStock {
			fmt.Println("Waiting for item")
			cond.Wait() // 대기 중일 때 뮤텍스를 자동으로 해제하고, 깨어날 때 뮤텍스를 다시 잡음
		}
		item.buy()
		mu.Unlock()
	}()

	// 아이템 재고 충전 및 대기 중인 고루틴 깨우기를 시뮬레이션
	time.Sleep(time.Second)
	mu.Lock()
	item.inStock = true
	mu.Unlock()
	cond.Signal() // 대기 중인 고루틴 중 하나를 깨움

	time.Sleep(time.Second) // 출력을 보기 위해 대기
}
