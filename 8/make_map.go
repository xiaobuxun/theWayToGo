package main

import "fmt"

func main() {
	mapOrigin := map[string]string {
		"a": "A",
		"b": "B",
		"c": "C",
	}

	mapTarget := mapOrigin

	mapTarget["a"] = "T"

	fmt.Println(mapOrigin)
	fmt.Println(mapTarget)
}