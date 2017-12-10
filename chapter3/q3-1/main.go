package main

import (
	"io"
	"log"
	"os"
)

func main() {
	old, err := os.Open("old.txt")
	if err != nil {
		log.Fatal(err)
	}

	new, err := os.Create("new.txt")
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(new, old)
	if err != nil {
		log.Fatal(err)
	}
}
