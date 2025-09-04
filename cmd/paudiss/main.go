package main

import (
	"log"
	"os"

	"github.com/kosper/pauvm/pkg/utils"
	"github.com/kosper/pauvm/internal/paudiss"
)

const Usage string = "Usage: paudiss.exe <bytecodeFile>"

func main() {
	var osArgs []string = os.Args[1:]
	var argsLen = len(osArgs)

	if argsLen < 1 {
		log.Fatal(Usage)
	}
	
	var filename string = osArgs[0]

	if utils.IsFileExtension(filename, "pau") == false {
		log.Fatalf("File %s is not of .pau extension", filename)
	}

	if err := paudiss.Dissassemble(filename); err != nil {
		log.Fatal(err)
	}

	return
}
