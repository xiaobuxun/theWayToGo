package main

import "fmt"

func main() {
	var ch int = '\u0041'
	var ch1 int = '\u0051'
	var ch2 int = '\u0061'

	fmt.Printf("%d-%d-%d\n", ch, ch1, ch2)
	fmt.Printf("%c-%c-%c\n", ch, ch1, ch2)
	fmt.Printf("%X-%X-%X\n", ch, ch1, ch2)
	fmt.Printf("%U-%U-%U\n", ch, ch1, ch2)
}
