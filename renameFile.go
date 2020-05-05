package main

import (
	"log"
	"os"
)

func main() {
	oldLocation := "./main"
	newLocation := "./fold/main"
	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		log.Fatal(err)
	}
}
