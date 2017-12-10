package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
)

func main() {
	out, err := os.Create("test.zip")
	if err != nil {
		log.Fatal(err)
	}

	zip := zip.NewWriter(out)
	defer zip.Close()

	writer, err := zip.Create("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	source, err := os.Open("main.go")
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(writer, source)
	if err != nil {
		log.Fatal(err)
	}
}
