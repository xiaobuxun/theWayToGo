package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "String to test upper and lower."
	fmt.Printf("The origin string is:%s\n", str)
	fmt.Printf("The upper string is:%s\n", strings.ToUpper(str))
	fmt.Printf("The lower string is:%s\n", strings.ToLower(str))
}
