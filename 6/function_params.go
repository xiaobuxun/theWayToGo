package main

import "fmt"

func add(a, b int) {
	fmt.Printf("%d + %d = %d\n", a, b, a + b)
}

func callback(m, n int, f func(int, int)) {
	f(m, n)
}

func main() {
	callback(1, 2, add)
}