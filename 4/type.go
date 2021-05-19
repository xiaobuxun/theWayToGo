package main

import "fmt"

func main() {
	type T int
	var a, b T = 3, 4
	c := a + b
	fmt.Printf("c value is:%d", c)
}
