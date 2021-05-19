package main

import (
	"fmt"
	"strings"
)

func main() {
	strField := "To test string field."
	resultField := strings.Fields(strField)

	for key, val := range resultField {
		fmt.Println("Fields result:")
		fmt.Printf("key:%d, value:%s\n", key, val)
	}

	strSplit := "To,test,string,split."
	resultSplit := strings.Split(strSplit, ",")
	for key, val := range resultSplit {
		fmt.Println("Split result:")
		fmt.Printf("key:%d, value:%s\n", key, val)
	}

	strJoin := strings.Join(resultSplit, ";")
	fmt.Printf("Join string:%s", strJoin)
}
