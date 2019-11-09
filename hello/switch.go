package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	x := 25
	switch x {
	case 10:
		fmt.Println("X is 10")
	case 25:
		fmt.Println("X is 25")
	default:
		fmt.Println("Fell on default case")
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	y := r1.Intn(100)
	switch {
	case y%2 == 0:
		fmt.Println("%i It's en even number", y)
	default:
		fmt.Println("%i It's an odd number", y)
	}
}
