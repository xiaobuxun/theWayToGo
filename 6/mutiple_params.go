package main

import "fmt"

func min(n ...int) int {
	if len(n) == 0 {
		return 0
	}
	res := n[0]
	for _, value := range n {
		if value < res {
			res = value
		}
	}
	return res
}

func main() {
	result := min(1,2,3,4,5)
	fmt.Printf("min is:%d\n", result)
	nums := []int{1, 2, 3, 4}
	result = min(nums...)
	fmt.Printf("min is:%d\n", result)
}