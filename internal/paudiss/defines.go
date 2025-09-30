package paudiss

import (
	"fmt"
	"os"
	"runtime"

	"github.com/kosper/pauvm/pkg/utils"
)

type PDFlags struct {
	filename string
}

func printUsage() {
	switch runtime.GOOS {
		case "windows": {
			var UsageWindows string = "Usage: paudiss.exe <bytecodeFile>"
			fmt.Println(UsageWindows)

			return
		}

		case "linux": {
			var UsageLinux string = "Usage: ./paudiss <bytecodeFile>"
			fmt.Println(UsageLinux)

			return
		}

		default: {
			var UsageLinux string = "Usage: ./paudiss <bytecodeFile>"
			fmt.Println(UsageLinux)

			return
		}
	}
}

func printHelp() {
	//Note: This is a multiline string, formatting is a little weird.
	var helpString string = 
`Paudiss is a program that dissassembles PauVM executable files.

Usage:
  paudiss <bytecodefile>
`
		fmt.Println(helpString)
}

func HandleConsoleArgs() (*PDFlags, error) {
	var flags PDFlags = PDFlags{
		filename: "",
	}

	var osArgs []string = os.Args[1:]
	var argsLen = len(osArgs)

	if argsLen < 1 {
		printUsage()
		os.Exit(0)
	}

	if osArgs[0] == "-help" {
		printHelp()
		os.Exit(0)
	}
	
	var filename string = osArgs[0]

	if err := utils.IsFileExtension(filename, ".pau"); err != nil {
		return &flags, err
	}

	flags.filename = filename

	return &flags, nil
}
