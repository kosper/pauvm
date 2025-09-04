package main

import (
	"github.com/kosper/pauvm/internal/pauvm"
	"github.com/kosper/pauvm/pkg/utils"

	"os"
	"log"
)

//TODO: Better error handling.
func main() {
	var args []string = os.Args[1:]

	if len(args) < 1 {
		log.Fatal("Usage: pauvm.exe <bytecodeFile>")
	}

	var filename string = args[0]

	if utils.IsFileExtension(filename, "pau") == false {
		log.Fatal("<bytecodeFile> should be of .pau extension")
	}

	var flags pauvm.VMFlags = pauvm.VMFlags{
		Trace: false,
	}

	if len(args) > 1 && args[1] == "-trace" {
		flags.Trace = true
	}

	var pauVM *pauvm.VM = pauvm.InitVM(&flags)

	if err := pauVM.LoadProgramFromFile(filename); err != nil {
		log.Fatal(err.Error())
	}

	pauVM.PrintProgram()

	var err error = pauVM.ExecuteProgram()

	if err != nil {
		log.Fatal(err.Error())
	}


	return
}
