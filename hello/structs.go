package main

import "fmt"

func main() {
	type Person struct {
		Name string
		Age  int
	}

	nanodano := &Person{Name: "Pepe", Age: 44}
	fmt.Println(nanodano)

	type Hacker struct {
		Person           Person
		FavoriteLanguage string
	}

	a_hacker := &Hacker{Person: *nanodano, FavoriteLanguage: "Go"}
	fmt.Println(a_hacker)
}
