package main

import "fmt"

func main() {
	from := []int{1,2,3}
	to := make([]int, 10)

	copyValue := copy(to, from)

	fmt.Printf("copy value is:%d\n", copyValue)

	to = append(to, 1,2,3,4,5)
	fmt.Printf("append value is:%d\n", to)
}