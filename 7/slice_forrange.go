package main

import "fmt"

func main() {
	seasons := []string{"one", "two", "three", "four"}

	for _,v := range seasons {
		fmt.Printf("value is:%s\n", v)
	}

	for k := range seasons {
		fmt.Printf("key is:%d\n", k)
	}
}