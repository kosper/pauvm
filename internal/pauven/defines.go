package pauven

import (
	"fmt"
	"os"
	"runtime"
	"errors"

	"github.com/kosper/pauvm/pkg/utils"
)

type CompilerFlags struct {
	MainFile string
	OutputName string
}

type Compiler struct {
	labels map[string]int32
	globalLabelOffset int32

	macros map[string]string

	sourceFiles []SourceFile

	CompilerFlags
}

type SourceFile struct {
	filename string
	content string
}

func printUsage() {
	switch runtime.GOOS {
		case "windows": {
			var UsageWindows string = "Usage: pauven.exe -f <filename> [-o <fileoutput>]"
			fmt.Println(UsageWindows)

			return
		}

		case "linux": {
			var UsageLinux string = "Usage: ./pauven.exe -f <filename> [-o <fileoutput>]"
			fmt.Println(UsageLinux)

			return
		}

		default: {
			var UsageLinux string = "Usage: ./pauven.exe -f <filename> [-o <fileoutput>]"
			fmt.Println(UsageLinux)

			return
		}
	}
}

func printHelp() {
	//Note: This is a multiline string, formattinf is weird.
	var helpString = 
`Pauven is a program that turns Pau bytecode into a PauVM ececutable

Usage:
  pauven -f <paufile> -o <output>
`
	fmt.Println(helpString)
}

func HandleConsoleArgs() (*CompilerFlags, error){
	var osArgs []string = os.Args[1:]
	var osArgsLen int = len(osArgs)

	var flags CompilerFlags = CompilerFlags{
		MainFile: "",
		OutputName: "out.pau",
	}

	if osArgs[0] == "-help" {
		printHelp()
		os.Exit(0)	
	}

	for i:= 0; i < osArgsLen; i++ {
		switch osArgs[i] {
			case "-o":
				//Note: -o flag requires filename.
				var nextIndex int = i + 1

				//Note: If there is no next flag return error.
				if nextIndex >= osArgsLen {
					var ferror string = fmt.Sprintf("Error: output name was not provided")
					return &flags, errors.New(ferror)
				}
				
				flags.OutputName = osArgs[nextIndex]

				if err := utils.IsFileExtension(flags.OutputName, ".pau"); err != nil {
					return &flags, err
				}

				//Note: We already handled next index.
				i++

				continue;
			case "-f": 
				//Note: -f flag requires filename.
				var nextIndex int = i + 1

				//Note: If there is no next flag return error.
				if nextIndex >= osArgsLen {
					var ferror string = fmt.Sprintf("Error: main file was not provided")
					return &flags, errors.New(ferror)
				}

				flags.MainFile = osArgs[nextIndex]

				if err := utils.IsFileExtension(flags.MainFile, ".pv"); err != nil {
					return &flags, err
				}

				//Note: We already handled next index.
				i++

				continue;
			default:
				var ferror string = fmt.Sprintf("Error: Flag %s does not exist", osArgs[i])
				return &flags, errors.New(ferror)
		}

	}

	return &flags, nil
}
