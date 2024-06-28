package main

import (
	"fmt"
	"reflect"
)

func main() {
	var string1 = "This is a string"
	discoverType(string1)

	discoverType(&string1)
	discoverType(1)
	discoverType(nil)
}

func discoverType(t any) {
	switch v := t.(type) {
	case string:
		fmt.Printf("String found: %s\n", v)
	case *string:
		fmt.Printf("String pointer found: %v\n", v)
	case int:
		fmt.Printf("Integer found: %d\n", v)
	default:
		theType := reflect.TypeOf(v)
		if theType == nil {
			fmt.Println("Type is nil")
		} else {
			fmt.Printf("Type not handled: %v\n", reflect.TypeOf(v))
		}
	}
}
