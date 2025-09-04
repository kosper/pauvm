package pauven

import (
	"log"
	"os"

	"github.com/kosper/pauvm/pkg/utils"
)

const Usage string = "Usage: pauven.exe -f <filename> [-o <fileoutput>]"

type SourceFile struct {
	content string
	formattedContent string

	labels map[string]int32
}

func HandleArgs(inputName *string, outputName *string) {
	var osArgs []string = os.Args[1:]

	if len(osArgs) < 2 || osArgs[0] != "-f" {
		log.Println(Usage)
		os.Exit(-1)	
	}

	*inputName = osArgs[1]
	*outputName = "out.pau"

	if utils.IsFileExtension(*inputName, "pv") == false {
		log.Println("File input name should be of extension .pv")
		os.Exit(-1)
	}

	if len(osArgs) > 3 {
		if osArgs[2] == "-o" {
			*outputName = osArgs[3]	

			if utils.IsFileExtension(*outputName, "pau") == false {
				log.Println("File output name should be of extension .pau")
				os.Exit(-1)
			}
		}
	}

	return
}
