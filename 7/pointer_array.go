package main

import "fmt"

func main() {
	var pointer_array *[3]int
	for i := 0; i <= 10; i++ {
		pointer_array = &[3]int{i, i * i, i * i * i}
		fmt.Printf("pointer is:%p\n", pointer_array)
	}
}