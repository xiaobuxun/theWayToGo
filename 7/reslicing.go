package main

import "fmt"

func main() {
	slice := make([]int, 0, 10)
	for i := 0; i < cap(slice); i++ {
		slice = slice[0 : i + 1]
		slice[i] = i
		fmt.Printf("The length of slice is:%d\n", len(slice))
	}

	for i := 0; i < len(slice); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice[i])
	}
}