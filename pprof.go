package main

import (
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
	"time"
)

func busyWork() {
	for i := 0; i < 1e3; i++ {
		time.Sleep(time.Millisecond)
	}
}

// cmd: go tool pprof cpu.pprof
func Pprof() {
	f, err := os.Create("cpu.pprof")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// CPU 프로파일링 시작
	if err := pprof.StartCPUProfile(f); err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()

	// 프로파일링할 작업
	busyWork()
}

// see: http://0.0.0.0:6060/debug/pprof
func WebPprof() {
	http.ListenAndServe("0.0.0.0:6060", nil)
}
