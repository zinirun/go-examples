package main

import "fmt"

func SomePanicFunction() {
	panic("This is a panic")
}

func RecoverUncaughtExceptions() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	SomePanicFunction()
}
