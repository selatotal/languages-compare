package main

import (
	"time"
)

func main() {

	go func() {
		println("Hello Routine")
	}()
	println("Hello Main")
	time.Sleep(time.Second)
}
