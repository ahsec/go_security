package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) SayMyName() {
	fmt.Println("My name is:", p.Name)
}
func (p Person) ChangeMyName(newName string) {
	p.Name = newName
}

func main() {
	pepe := Person{Name: "Pepe"}
	pepe.SayMyName()
	pepe.ChangeMyName("Antonio")
	pepe.SayMyName()
}
