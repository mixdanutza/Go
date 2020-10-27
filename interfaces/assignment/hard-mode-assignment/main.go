package main

import (
	"fmt"
	"io"
	"os"
)

type logWriter struct{}

func main() {
	filePath := os.Args[1]
	fileContent, error := os.Open(filePath)
	lw := logWriter{}

	if error != nil {
		fmt.Println("Error: ", error)
		os.Exit(1)
	}

	io.Copy(lw, fileContent)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return (len(bs)), nil
}
