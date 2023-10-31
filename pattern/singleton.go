package main

import (
	"sync"
)

type singleton struct {
}

var instance *singleton
var once1 sync.Once

// GetInstance는 싱글톤 객체의 유일한 인스턴스를 반환
func GetInstance() *singleton {
	// 여러 고루틴에서 동시에 실행하더라도 한번만 실행
	once1.Do(func() {
		instance = &singleton{}
	})
	return instance
}
