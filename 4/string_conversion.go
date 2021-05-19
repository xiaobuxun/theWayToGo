package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "666"
	var an int
	var newStr string

	fmt.Printf("int size is:%d\n", strconv.IntSize)
	an, _ = strconv.Atoi(str)
	fmt.Printf("The ingeger is:%d\n", an)
	an += 5
	newStr = strconv.Itoa(an)
	fmt.Printf("The new string is:%s\n", newStr)
}
