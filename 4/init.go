package main

import (
	"fmt"
	"math"
)

var Pi float64

func init() {
	Pi = math.Tan(1) * 4
}

func main() {
	fmt.Printf("Pi value is:%v", Pi)
}
