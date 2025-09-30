package main

import (
	"log"

	"github.com/kosper/pauvm/internal/pauven"
)

func main() {
	var mainFile pauven.SourceFile

	flags, err := pauven.HandleConsoleArgs()

	if err != nil {
		log.Fatal(err)
	}

	if flags.MainFile == "" {
		log.Fatal("Main file was not provided.")
	}

	if err := mainFile.Read(flags.MainFile); err != nil {
		log.Fatal(err)
	}

	//TODO: Accept flags.
	var compiler = pauven.CompilerInit(flags.OutputName)

	if err := compiler.Preprocess(&mainFile); err != nil {
		log.Fatal(err)
	}

	if err := compiler.Compile(); err != nil {
		log.Fatal(err)
	}

	return
}
