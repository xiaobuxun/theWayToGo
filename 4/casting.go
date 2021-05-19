package main

import "fmt"

func main() {
	var number16 int16 = 34
	var number32 int32

	number32 = int32(number16)

	fmt.Printf("int16 value is:%d\n", number16)
	fmt.Printf("int32 value is:%d\n", number32)
}
