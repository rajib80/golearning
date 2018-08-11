package main

import (
	"fmt"
	"os"
)

func main() {
	textFile := createFile("/Users/rajib/Temp/TextFileForTestingDefer.txt")
	defer closeFile(textFile)
	writeFile(textFile)
}

func createFile(fileNameWithPath string) *os.File {
	fmt.Println("Creating file ...")
	file, error := os.Create(fileNameWithPath)
	if error != nil {
		panic(error)
	}
	return file
}

func writeFile(inputFile *os.File) {
	fmt.Println("Writing file ...")
	fmt.Fprintln(inputFile, "This file is created to test defer")
}

func closeFile(fileToBeClosed *os.File) {
	fmt.Println("Closing file ...")
	fileToBeClosed.Close()
}
