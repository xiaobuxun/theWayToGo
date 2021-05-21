package main

import "fmt"

type twoInt struct {
	a int
	b int
}

func main() {
	two2 := twoInt{1, 2}
	fmt.Printf("The sum is:%d\n", two2.addItem())
	fmt.Printf("The sum add param is:%d\n", two2.addParam(3))
}

func (n * twoInt) addItem() int {
	return n.a + n.b
}

func (n * twoInt) addParam(add int) int {
	return n.a + n.b + add
}