package pauven

import (
	"bufio"
	"strings"
	"log"
	"errors"
	"fmt"
)

func (compiler *Compiler)Preprocess(file *SourceFile) error {
	var scanner *bufio.Scanner = bufio.NewScanner(strings.NewReader(file.content))
	var linecnt = 0

	log.Printf("Preprocessing file %s...\n", file.filename)

	/*
	* Note: For every line, trim left-right, if the line is empry or contains a comment ignore.
	* If the last character of the line is the : character, it represents a label-fucntion, we add the 
	* label to the labels map along with the number of the current non empty-non commented line(represents the ip).
	*/
	for scanner.Scan() {
		linecnt++
		var line string = scanner.Text()

		line = strings.TrimLeft(line, " \t");
		line = strings.TrimRight(line, " \t");

		if line == "" || line[0] == '#' {
			continue
		}

		//Note: Handle maccros.
		if line[0] == '!' {
			err := compiler.handleMacro(line)

			if err != nil {
				var errstr string = fmt.Sprintf("%s:%d: %s", file.filename, linecnt, err) 
				return errors.New(errstr)
			}

			continue
		}

		if line[len(line) - 1] == ':' {
			err := compiler.handleLabel(line)

			if err != nil {
				var errstr string = fmt.Sprintf("%s:%d: %s", file.filename, linecnt, err)
				return errors.New(errstr)
			}

			continue
		}

		compiler.globalLabelOffset++
	}

	compiler.sourceFiles = append(compiler.sourceFiles, *file)

	log.Printf("File %v preprocessed\n", file.filename)

	return nil
}

func (compiler *Compiler)handleMacro(line string) error {
	var macro string = line[1:]

	//Note: Define flag.
	//TODO: Change from def to define.
	if macro[0:4] == "def " {
		/*
		 * Note: after the def keyword we set the expression
		 * so every time the expression begins in the 3rd position of the line.
		 * The definition should be one word, thus we search the expression until an
		 * empty character is found. Before the empty character is the definition name.
		 * And after the empty character is the defined expression.
		 */
		var expression string = macro[3:]
		var define string //Note: Definition
		var token string //Note: Defined expression.
		var endIndex = 0

		expression = strings.TrimLeft(expression, " \t")

		for i := range(len(expression) + 1) {
			if expression[i] == ' ' {
				token = expression[(i+1):]
				token = strings.TrimLeft(string(token), " \t")
				break;
			}

			endIndex++
		}

		define = expression[0: endIndex]

		if len(token) < 1 || len(expression) < 1 {
			var errstr = fmt.Sprintf("macro (%s) definition is wrong!", macro) 
			return errors.New(errstr)
		}

		//Note: Macro is already defined.
		if _, ok := compiler.macros[define]; ok == true {
			var errstr = fmt.Sprintf("macro (%s) is already defined!", macro) 
			return errors.New(errstr)
		}

		if val, ok := compiler.macros[token]; ok == true {
			compiler.macros[define] = val
		} else {
			compiler.macros[define] = token
		}
	} else if macro[0:8] == "include " {
		//Note: If a file is included, preprocess recursively.
		var tokens = strings.Fields(macro)

		if len(tokens) < 2 {
			return errors.New("include file must be specified!")
		}
		
		var filename string = tokens[1]

		//Note: Check if file is already parsed.
		for i := range(len(compiler.sourceFiles)) {
			if compiler.sourceFiles[i].filename == filename {
				return nil
			}
		}

		var includeFile SourceFile	

		if err := includeFile.Read(filename); err != nil {
			return err
		}

		if err := compiler.Preprocess(&includeFile); err != nil {
			return err
		}
	}

	return nil
}

func (compiler *Compiler)handleLabel(line string) error {
	//Note: Trim every empty character and : character from the end of the line, so we can keep only the label.
	var label string = strings.TrimRight(line, " :\t")

	//Note: If label is already defines, return error.
	if _, ok := compiler.labels[label]; ok == true {
		var errstr string = fmt.Sprintf("Label is %s already specified!", label)
		return errors.New(errstr)
	}

	//Note: Update global label offset with the offset of the label.
	compiler.labels[label] = compiler.globalLabelOffset

	return nil
}
