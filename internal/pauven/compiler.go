package pauven

import (
	"os"
	"log"
	"fmt"
	"strings"
	"errors"
	"encoding/binary"
	"unicode"
	"bufio"
	"strconv"
	
	"github.com/kosper/pauvm/pkg/isa"
)


//NOTE: Reads entire file into memory.
func (file *SourceFile)Read(filename string) error {
	b, err := os.ReadFile(filename)
	
	if err != nil {
		return err
	}

	file.content = string(b)

	return nil
}

//NOTE: Translates code into opcodes and values(Should be used after pre-processor)
func (file *SourceFile)Compile(outputName string) error {
	var scanner *bufio.Scanner = bufio.NewScanner(strings.NewReader(file.formattedContent))

	log.Println("Compiling File...")

	fileoutput, err := os.Create(outputName)

	if err != nil {
		return err
	}

	defer fileoutput.Close()

	/*
	* NOTE: For every line, trim left-right(It is not needed because we trimmed the in the preprocessor, but i do what i want), we translate the
	* opcode and we write to a binary file the translated byte. If the opcode contains a value then we also write the int32 value in little endian format.
	* If the value is a label then we search the labels map and replace the label with the ip.
	*/
	for scanner.Scan() {
		var line string = scanner.Text()

		var tokens []string = strings.Fields(line)

		opCode, ok := isa.StringToInst[tokens[0]]

		if ok == false {
			return errors.New(fmt.Sprintf("Uknown Instruction %s", tokens[0]))
		}

		if err := binary.Write(fileoutput, binary.LittleEndian, byte(opCode)); err != nil {
			return err
		}

		if isa.InstHasOperand[opCode] == true {
			var val int32 = 0

			if len(tokens) < 2 {
				return errors.New(fmt.Sprintf("Opcode %s requires operand", isa.InstToString[opCode]))
			}

			var containsLabel = (opCode == isa.INST_JMP || opCode == isa.INST_JMPZ || opCode == isa.INST_CALL)

			if  containsLabel && (unicode.IsDigit(rune(tokens[1][0])) == false) {
				val, ok = file.labels[tokens[1]]

				if ok == false {
					log.Fatal(fmt.Sprintf("Label %s does not exist", tokens[1]))
				}
			} else {
				val64, _ := strconv.ParseInt(tokens[1], 10, 32)
				val = int32(val64)
			}

			if err := binary.Write(fileoutput, binary.LittleEndian, int32(val)); err != nil {
				return err
			}
		}
	}

	log.Println("File compiled successfully")

	return nil
}
