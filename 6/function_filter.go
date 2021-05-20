package main

import "fmt"

type validate func(int) bool

func isOdd(n int) bool {
	if (n % 2 == 0) {
		return false
	}
	return true
}

func isEven(n int) bool {
	if (n % 2 == 0) {
		return true
	}
	return true
}

func filter(origin []int, test validate) [] int {
	var result []int
	for _,value := range origin {
		if (test(value)) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	slice := []int{1,2,3,4,5}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd)
	fmt.Println("odd items are:", odd)
	even := filter(slice, isEven)
	fmt.Println("even items are:", even)
}