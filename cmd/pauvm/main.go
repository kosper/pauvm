package main

import (
	"github.com/kosper/pauvm/internal/pauvm"
	"log"
)

func main() {
	flags, err := pauvm.HandleConsoleArgs()

	if err != nil {
		log.Fatal(err.Error())
	}

	var pauVM *pauvm.VM = pauvm.InitVM(flags)

	if err := pauVM.LoadProgramFromFile(flags.Filename); err != nil {
		log.Fatal(err.Error())
	}

	pauVM.PrintProgram()

	err = pauVM.ExecuteProgram()

	if err != nil {
		log.Fatal(err.Error())
	}

	return
}
