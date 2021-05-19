package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = "hello, how it is going on?"

	fmt.Printf("The number of 'l' in '%s' is:%d\n", str, strings.Count(str, "l"))
}
