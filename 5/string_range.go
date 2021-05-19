package main

import "fmt"

func main() {
	str := "English string"
	strChinese := "中文字符串"

	fmt.Printf("English string length:%d\n", len(str))
	fmt.Printf("Chinese string length:%d\n", len(strChinese))

	for pos, char := range str {
		fmt.Printf("Charactor on position %d is:%c\n", pos, char)
	}
	for pos, char := range strChinese {
		fmt.Printf("Charactor on position %d is:%c\n", pos, char)
	}
}
