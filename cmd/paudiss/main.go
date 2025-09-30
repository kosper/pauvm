package main

import (
	"log"
	"github.com/kosper/pauvm/internal/paudiss"
)

func main() {
	flags, err := paudiss.HandleConsoleArgs()

	if err != nil {
		log.Fatal(err)
	}

	if err := paudiss.Dissassemble(flags); err != nil {
		log.Fatal(err)
	}

	return
}
