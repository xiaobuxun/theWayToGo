package main

import "fmt"

func main() {
	num := 1
	fmt.Printf("Number value is:%d\n", num)
	num, err := function1(num)
	if err != nil {
		fmt.Printf("Error when run function1:%s\n", err)
	}

	fmt.Printf("Number after function1 value is:%d\n", num)
}

func function1(origin int) (int, error) {
	origin = origin * 100
	return origin, nil
}
