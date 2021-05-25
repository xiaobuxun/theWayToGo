package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

type address struct {
	country string
	city string
	street string
	mobile string
}

type vcard struct {
	firstname string
	lastname string
	address []*address
	remark string
}

func main() {
	address1 := &address{
		"cn",
		"beijing",
		"haidian",
		"186...",
	}
	address2 := &address{
		"cn",
		"beijing",
		"dongcheng",
		"186...",
	}
	vc := vcard{
		"li",
		"gang",
		[]*address{address1, address2},
		"remark",
	}

	file,_ := os.OpenFile("vcard.god", os.O_CREATE | os.O_WRONLY | os.O_RDWR, 0666)

	defer file.Close()

	enc := gob.NewEncoder(file)
	err := enc.Encode(vc)

	if err != nil {
		fmt.Printf("error in encode gob:%v\n", err)
	}
}