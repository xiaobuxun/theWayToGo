package main

import (
	"fmt"
	"reflect"
)

type tag struct {
	Name string `json:"n" xml:"n"`
	Sex string `json:"s" xmc:"s"`
}

func main() {
	testTag := tag{"li","man"}
	types := reflect.TypeOf(testTag)
	
	for _,v := range []string{"Name","Sex"} {
		field, found := types.FieldByName(v)
		if !found {
			continue
		}

		fmt.Printf("file name:%s\n", v)
		fmt.Printf("whole tag is:%s\n", field.Tag)
		fmt.Printf("value of json:%s\n", field.Tag.Get("json"))
	}
}