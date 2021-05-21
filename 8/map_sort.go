package main

import (
	"fmt"
	"sort"
)

func main() {
	barVal := map[string]int{
		"alpha": 34,
		"bravo": 56,
		"charlie": 23,
		"delta": 87,
		"echo": 56,
		"foxtrot": 12,
		"golf": 34,
		"hotel": 16,
		"indio": 87,
		"juliet": 65,
		"kilo": 43,
		"lima": 98,
	}

	fmt.Println("unsorted:")

	for k,v := range barVal {
		fmt.Printf("key:%s, value:%d\n", k, v)
	}

	var keys []string

	for k := range barVal {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	fmt.Println("sorted:")

	for _,v := range keys {
		fmt.Printf("key:%s, value:%d\n", v, barVal[v])
	}
}