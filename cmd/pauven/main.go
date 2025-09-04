package main

import (
	"log"

	"github.com/kosper/pauvm/internal/pauven"
)

func main() {
	var inputName string
	var outputName string
	var mainFile pauven.SourceFile

	//TODO: Flags
	pauven.HandleArgs(&inputName, &outputName)
	
	if err := mainFile.Read(inputName); err != nil {
		log.Fatal(err)
	}

	if err := mainFile.Preprocess(); err != nil {
		log.Fatal(err)
	}

	if err := mainFile.Compile(outputName); err != nil {
		log.Fatal(err)
	}

	return
}
