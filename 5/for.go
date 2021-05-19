package main

import "fmt"

func main() {
L:
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			if j == 3 {
				continue L
			}
			fmt.Printf("i is:%d and j is:%d\n", i, j)
		}
	}
}
