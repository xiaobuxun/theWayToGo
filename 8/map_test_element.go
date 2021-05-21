package main

import "fmt"

func main() {
	origin := map[string]int {
		"a": 1,
		"b": 2,
		"c": 3,
	}

	value, isPresent := origin["a"]
	if (isPresent) {
		fmt.Printf("The value of key(%s) in origin is:%d\n", "a", value)
	} else {
		fmt.Println("There is no key(a) in origin.")
	}

	delete(origin, "a")
	value, isPresent = origin["a"]
	if (isPresent) {
		fmt.Printf("The value of key(%s) in origin is:%d\n", "a", value)
	} else {
		fmt.Println("There is no key(a) in origin.")
	}

}