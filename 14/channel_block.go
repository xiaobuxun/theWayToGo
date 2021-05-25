package main

import (
	"fmt"
	"time"
)

func pump(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func suck(ch chan int) {
	for {
		x := <-ch
		fmt.Printf("channel value is:%d\n", x)
	}
}

func main() {
	ch1 := make(chan int)
	go pump(ch1)
	// go suck(ch1)
	time.Sleep(3e9)
}