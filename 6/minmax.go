package main

import "fmt"

func minmax(m, n int)(min int, max int) {
	if m < n {
		min = m
		max = n
	} else {
		min = n
		max = m
	}
	return
}

func main() {
	var min, max int
	min, max = minmax(1, 10)
	fmt.Printf("min is:%d\nmax is:%d\n", min, max)
}