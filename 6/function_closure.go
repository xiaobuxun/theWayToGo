package main

import "fmt"

func closure() func(int) int {
	var x int
	return func(delta int) int {
		fmt.Printf("x:%d\ndelta:%d\n", x, delta)
		x += delta
		return x
	}
}

func fibonacci() func(int) int {
	var fibonaccis []int
	return func(n int) int {
		var result int
		switch n {
		case 0:
			result = 1
		case 1:
			result = 1
		default:
			result = fibonaccis[n-1] + fibonaccis[n-2]
		}
		fibonaccis = append(fibonaccis, result)
		return result
	}
}

func main() {
	// var f = closure()
	// fmt.Print(f(1), "-")
	// fmt.Printf("result:%d\n", f(20))
	// fmt.Printf("result:%d\n", f(300))
	var f = fibonacci()
	for i := 0; i <= 10; i++ {
		fmt.Printf("fibonacci(%d) result is:%d\n", i, f(i))
	}
}
