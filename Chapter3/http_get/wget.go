package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	outFname = "output.html"
	url = "http://devdungeon.com/archive"
)

func main(){
	newFile, err := os.Create(outFname)
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	writtenBytes, err := io.Copy(newFile, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Download of: %v completed\n", url)
	fmt.Printf("File: %v  saved. Bytesize: %v\n", outFname, writtenBytes)
}