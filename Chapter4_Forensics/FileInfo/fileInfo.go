package main

import (
	"fmt"
	"log"
	"os"
)

var (
	fname = "test.txt"
	err error
	fileInfo os.FileInfo
)

func main(){
	fileInfo, err := os.Stat(fname)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Filename: ", fileInfo.Name())
	fmt.Println("Size in Bytes: ", fileInfo.Size())
	fmt.Println("Permissions: ", fileInfo.Mode())
	fmt.Println("Last modified: ", fileInfo.ModTime())
	fmt.Println("Is directory: ", fileInfo.IsDir())
	fmt.Println("System Interface: ", fileInfo.Sys())
}