package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileName := "test.txt"
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		log.Fatal("Error obtaining info on file: ", fileName)
	}
	fmt.Println(fileInfo)
	fmt.Println("Name: ", fileInfo.Name())
}
