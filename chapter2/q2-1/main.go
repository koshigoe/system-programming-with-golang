package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("out.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(f, "%%d: %d, %%f: %f, %%s: %s\n", 1, 1.0, "1")
}
