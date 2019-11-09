package main

import "fmt"

func main() {
	slice := []int{2, 4, 6, 7}
	for key, value := range slice {
		fmt.Println(key, value)
	}

	natoMap := map[string]string{
		"A": "Alpha",
		"B": "Bravo",
		"C": "Charlie",
		"D": "Delta",
		"E": "Echo",
		"F": "Foxtrot",
		"G": "Golf",
		"H": "Hotel",
	}
	for key, value := range natoMap {
		fmt.Println(key, value)
	}
}
