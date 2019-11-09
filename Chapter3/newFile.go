package main

import (
	"log"
	"os"
)

func main() {
	fileName := "test.txt"
	newFile, err := os.Create(fileName)
	defer newFile.Close()
	if err != nil {
		log.Fatal("Can't create file: ", fileName)
	}
	log.Println(newFile)
}
