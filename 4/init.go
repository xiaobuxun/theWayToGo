package main

import (
	"fmt"
	"math"
)

var pi float64

func init() {
	pi = math.Tan(1) * 4
}

func main() {
	fmt.Printf("pi value is:%v", pi)
}
