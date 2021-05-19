package main

import (
	"fmt"
	"time"
)

func main() {
	var week time.Duration = 3600 * 24 * 7 * 1e9
	var t = time.Now()
	//fmt.Println(t)
	fmt.Printf("%02d-%02d-%02d\n", t.Year(), t.Month(), t.Day())
	//t = time.Now().UTC()
	//fmt.Println(t)
	//fmt.Println(time.Now())
	fmt.Println(t.Add(week))
	fmt.Println(t.Format(time.RFC822))
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format("02 January 2006 15:05"))
	format := t.Format("20060102150405")
	fmt.Println(t, "=>", format)
}
