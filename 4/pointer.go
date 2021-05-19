package main

import "fmt"

func main() {
	param := 5
	fmt.Printf("The location in memory of number('%d') is:%p", param, &param)
}
