package main

import "fmt"

type one struct{}

func (n *one)takePicture() string {
	return "picture"
}

type two struct{}

func (n *two)call() string {
	return "ring"
}

type three struct{
	one
	two
}

func main() {
	Three := new(three)
	fmt.Printf("call one takePicture:%s\n", Three.takePicture())
	fmt.Printf("call two call:%s\n", Three.call())
}
