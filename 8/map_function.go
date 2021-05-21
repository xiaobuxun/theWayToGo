package main

import "fmt"

func main() {
	mapTarget := make(map[int]int)
	mapFunc := map[int]func() int {
		1: func() int {
			return 10
		},
		2: func() int {
			return 20
		},
		3: func() int {
			return 30
		},
	}

	for k,v := range mapFunc {
		mapTarget[k] = v()
	}

	fmt.Println(mapTarget)
}