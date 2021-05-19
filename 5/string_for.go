package main

import "fmt"

func main() {
	str := "string to test for."
	fmt.Printf("string length is:%d\n", len(str))
	for i := 0; i < len(str); i++ {
		fmt.Printf("Character on position %d is:%c\n", i, str[i])
	}
	strChinese := "for循环中文测试"
	fmt.Printf("string length is:%d\n", len(strChinese))
	for i := 0; i < len(strChinese); i++ {
		fmt.Printf("Character on position %d is:%c\n", i, strChinese[i])
	}
}
