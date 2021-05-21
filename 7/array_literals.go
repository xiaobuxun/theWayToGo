package main

import "fmt"

func main() {
	normal := [5]int{1,2,3,4,5}
	lazy := []int{6,7,8,9,10}
	keyValue := []string{1: "one", 2: "two", 3: "three"}

	for _,v := range normal {
		fmt.Printf("normal value is:%d\n", v)
	}
	for _,v := range lazy {
		fmt.Printf("lazy value is:%d\n", v)
	}
	for k,v := range keyValue {
		fmt.Printf("key %d value is:%s\n", k, v)
	}
}