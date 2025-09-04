package pauven

import (
	"bufio"
	"strings"
	"log"
	"errors"
	"fmt"
)

func (file *SourceFile)Preprocess() error {
	var scanner *bufio.Scanner = bufio.NewScanner(strings.NewReader(file.content))
	var ip int32 = 0
	var sb strings.Builder 

	file.labels = make(map[string]int32)

	log.Println("Preprocessing file...")

	/*
	* NOTE: For every line, trim left-right, if the line is empry or contains a comment ignore.
	* If the last character of the line is the : character, it represents a label-fucntion, we add the 
	* label to the labels map along with the number of the current non empty-non commented line(represents the ip).
	*/
	for scanner.Scan() {
		var line string = scanner.Text()

		line = strings.TrimLeft(line, " \t");
		line = strings.TrimRight(line, " \t");

		if line == "" || line[0] == '#' {
			continue
		}

		if line[len(line) - 1] == ':' {
			var label string = strings.TrimRight(line, " :\t")

			if _, ok := file.labels[label]; ok == true {
				return errors.New(fmt.Sprintf("Label %s already exists!", label))
			}

			file.labels[label] = ip

			continue
		}

		sb.WriteString((line + "\n"))
		ip++
	}

	file.formattedContent = sb.String()

	log.Println("File preprocessed")

	return nil
}
