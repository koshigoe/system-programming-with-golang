package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	w := csv.NewWriter(os.Stdout)
	if err := w.Write([]string{"a", "b", "c"}); err != nil {
		log.Fatal(err)
	}

	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
