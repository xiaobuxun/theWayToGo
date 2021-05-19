package main

import (
	"fmt"
	"time"
)

var fibonaccies [10]uint64

func main() {
	var result uint64 = 0

	start := time.Now()

	for i := 1; i < 10; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is:%d\n", i, result)
	}

	end := time.Now()

	delta := end.Sub(start)
	fmt.Printf("Caculation took this amont of time:%s\n", delta)

}

func fibonacci(n int) (result uint64) {
	if fibonaccies[n] != 0 {
		result = fibonaccies[n]
		return
	}
	if n <= 1 {
		result = 1
	} else {
		result = fibonacci(n-1) + fibonacci(n-2)
	}
	fibonaccies[n] = result
	return
}
