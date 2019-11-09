package main

import (
	"fmt"
	"reflect"
)

func main() {
	type Person struct {
		Name string
		age  int
	}
	type Doctor struct {
		person     Person
		speciality string
	}
	nanoDano := Person{
		Name: "Dr Nano",
		age:  99,
	}
	drNano := Doctor{
		person:     nanoDano,
		speciality: "Hacking",
	}
	fmt.Println(reflect.TypeOf(nanoDano))
	fmt.Println(nanoDano)

	fmt.Println(reflect.TypeOf(drNano))
	fmt.Println(drNano)
}
