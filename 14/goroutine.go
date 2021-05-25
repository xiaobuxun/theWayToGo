package main

import (
	"fmt"
	"time"
)

func pump1(ch chan int) {
	for i:=0;i<=10;i++ {
		ch <- i * 2
	}
}

func pump2(ch chan int) {
	for i:=0;i<=10;i++ {
		ch <- i +5
	}
}

func suck(ch1,ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Printf("receive from channel 1:%d\n", v)
		case v := <-ch2:
			fmt.Printf("receive from channer 2:%d\n", v)
		}
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go pump1(ch1)
	go pump2(ch2)
	go suck(ch1,ch2)
	time.Sleep(2e9)
}