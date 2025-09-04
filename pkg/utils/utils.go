package utils 

import (
	"strings"
)

func IsFileExtension(filename string, extension string) bool {
	var substr []string = strings.SplitAfter(filename, ".")
	var arrLen = len(substr)

	if arrLen > 1 && substr[arrLen - 1] == extension {
		return true
	}

	return false
}
