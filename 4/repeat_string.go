package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "String fo test string repeat."
	fmt.Printf("repeat string:%s", strings.Repeat(str, 3))
}
