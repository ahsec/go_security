package main

import (
	"fmt"
	"reflect"
)

func main() {
	myInt := 43
	myIntPointer := &myInt

	fmt.Println(reflect.TypeOf(myIntPointer))
	fmt.Println(myIntPointer)
	fmt.Println(*myIntPointer)
	fmt.Println(myInt)
	*myIntPointer = 7
	fmt.Println(myIntPointer)
	fmt.Println(*myIntPointer)
	fmt.Println(myInt)
}
