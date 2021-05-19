package main

import "fmt"

var localA string = "G"

func main() {
	localN()
	localM()
	localN()
}

func localN() {
	fmt.Printf(" a is :%v\n", localA)
}

func localM() {
	localA = "O"
	fmt.Printf(" a is :%v\n", localA)
}
