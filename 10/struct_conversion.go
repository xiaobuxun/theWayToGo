package main

import "fmt"

type number struct {
	f float32
}

type number2 number

func main() {
	a := number{5.0}
	b := number2{5.0}
	c := number(b)

	fmt.Printf("a:%v\nb:%v\nc:%v\n", a,b,c)
}