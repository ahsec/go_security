// Package compression uses library gzip.
package main 

import (
	"compress/gzip"
	"strings"
	"log"
	"io"
	"os"
)

func compress(filename string) error {
	// Create .gz file to write to
	outputFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	// Create a gzip writter on top of the file writter
	gzipWritter := gzip.NewWriter(outputFile)
	defer gzipWritter.Close()

	_, err = gzipWritter.Write([]byte("Gophers rule\n"))
	if err != nil {
		return err
	}
	log.Printf("Compressed data saved to file: %v\n", filename)
	return nil
}

func decompress(filename string) error {
	if _, err := os.Stat(filename); err != nil {
		return err
	}
	gzipFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	gzipReader, err := gzip.NewReader(gzipFile)
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	decomFilename := returnFilename(filename)
	outFileWritter, err := os.Create(decomFilename)
	if err != nil {
		return err
	}
	defer outFileWritter.Close()

	if _, err := io.Copy(outFileWritter, gzipReader); err != nil {
		return err
	}
	log.Printf("File %v decompressed into: %v", filename, decomFilename)
	return nil
}

func returnFilename(filename string) string{
	// When given a string with a prefix .gz, returns the string without the prefix
	// When given a string without the prefix .gz, returns the string with the .gz prefix
	// The idea is to use this function for compressing or decompressing files
	if strings.HasSuffix(filename, ".gz") {
		return strings.Trim(filename, ".gz")
	}
	return strings.Join([]string{filename, ".gz"}, "")
}

func main() {
	filename := "test.txt.gz"
	if err := compress(filename); err != nil {
		log.Fatal(err)
	}
	if err := decompress(filename); err != nil {
		log.Fatal(err)
	}
}