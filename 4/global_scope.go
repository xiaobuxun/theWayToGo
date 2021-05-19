package main

import "fmt"

var globalA = 'G'

func main() {
	globalN()
	globalM()
	globalN()
}

func globalN() {
	fmt.Printf(" a is :%d\n", globalA)
}

func globalM() {
	globalA = 'O'
	fmt.Printf(" a is :%d\n", globalA)
}
