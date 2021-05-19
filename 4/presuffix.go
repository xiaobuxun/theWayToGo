package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "This is a testing for string suffix"

	fmt.Printf("Is string '%s' has suffix %s:%t", str, "fix", strings.HasSuffix(str, "test"))
}
