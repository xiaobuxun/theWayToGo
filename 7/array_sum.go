package main

import "fmt"

func main() {
	slice := []float64{9,0.9,0.1,999.877}
	var sum float64

	for _,v := range slice {
		sum += v
	}

	fmt.Println(sum)
}