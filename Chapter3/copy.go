package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	srcFname := flag.String("src", "", "File to copy from")
	dstFname := flag.String("dst", "", "File to copy to")
	flag.Parse()

	srcFile, err := os.Open(*srcFname)
	if err != nil {
		log.Fatal("Error opening file: ", *srcFname)
	}
	defer srcFile.Close()

	dstFile, err := os.Create(*dstFname)
	if err != nil {
		log.Fatal("Error creating file: ", *srcFname)
	}
	defer dstFile.Close()

	cpBytes, err := io.Copy(dstFile, srcFile)
	if err != nil {
		log.Fatal("Error copying file contents: ", err)
	}
	fmt.Println("Bytes copied: ", cpBytes)

	err = dstFile.Sync()
	if err != nil {
		log.Fatal("Error syncing new file")
	}
}
