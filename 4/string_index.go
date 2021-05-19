package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hi, how is going on, is everything fine?"

	fmt.Printf("The position of \"Hi\" is:%d\n", strings.Index(str, "Hi"))
	fmt.Printf("The position of 'going' is:%d\n", strings.Index(str, "going"))
}
