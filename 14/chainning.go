package main

import (
	"flag"
	"fmt"
)

func f(left, right chan int) {
	left <- 1 + <- right
}

func main() {
	flag.Parse()
	leftmost := make(chan int)
	ngoroutine := flag.Int("n", 100, "how many")

	var left,right chan int = nil, leftmost

	for i := 0; i < *ngoroutine; i++ {
		left, right = right, make(chan int)
		go f(left, right)
	}
	right <- 0
	x := <- leftmost
	fmt.Println(x)
}