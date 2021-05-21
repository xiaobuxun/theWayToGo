package main

import "fmt"

func main() {
	items := make([]map[int]int, 5)

	for i := range items {
		items[i] = make(map[int]int)
		items[i][i] = i + 1
	}

	fmt.Println(items)

	for _,v := range items {
		v = make(map[int]int)
		v[1] = 1
	}

	fmt.Println(items)

}