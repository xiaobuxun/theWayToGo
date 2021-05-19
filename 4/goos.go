package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Printf("Go runtime goos is:%s\n", runtime.GOOS)
	fmt.Printf("Go env path is:%s\n", os.Getenv("PATH"))
}
