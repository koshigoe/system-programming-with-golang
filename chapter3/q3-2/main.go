package main

import (
	"crypto/rand"
	"io"
	"log"
	"os"
)

func main() {
	out, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.CopyN(out, rand.Reader, 1024)
	if err != nil {
		log.Fatal(err)
	}
}
