package main

import (
	"fmt"
	"log"
	"bufio"
	"os"
)

var (
	fname = "test.txt"
)

func getCountWords(fname string) []string {
	var allWords []string 
	// Open file and create scanner on top of it
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal("Error opening file. Error: ", err)
	}
	// By default, scanner will scan for words
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for {
	  success := scanner.Scan()
  	  // Either there's an error or we reached EoF
	  if success == false {
	  	if err := scanner.Err(); err != nil {
			  log.Fatal("Error reading file.  Error: ", err)
		}
		// Reached end of file
		break
	  }
	  allWords = append(allWords, scanner.Text())
	}
	return allWords 
}

func main(){
	allWords:= getCountWords(fname)
	for i, word := range allWords{
		fmt.Printf("Word %v : %v\n", i+1, word)
	}
}
