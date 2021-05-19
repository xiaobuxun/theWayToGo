package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("第%d次随机得到的整数是:%d\n", i, rand.Intn(8))
	}
}
