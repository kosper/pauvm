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
	"github.com/kosper/pauvm/pkg/utils"
)


func CompilerInit(outputName string) *Compiler {
	var compiler Compiler

	compiler.labels = make(map[string]int32)
	compiler.macros = make(map[string]string)

	compiler.globalLabelOffset = 1 //Note: Due to CALL main.

	compiler.sourceFiles = make([]SourceFile, 0, 10)
	compiler.OutputName = outputName

	return &compiler
}

//Note: Reads entire file into memory.
func (file *SourceFile)Read(filename string) error {
	b, err := os.ReadFile(filename)
	
	file.filename = filename

	if err != nil {
		return err
	}

	file.content = string(b)

	return nil
}

//Note: Translates code into opcodes and values(Should be used after pre-processor)
func (compiler *Compiler)Compile() error {
	fileoutput, err := os.Create(compiler.OutputName)

	if err != nil {
		return err
	}

	defer fileoutput.Close()

	//Note: Call main, first things first.
	val, ok := compiler.labels["main"]

	if ok == false {
		return errors.New("main is not defined, must be defined!")
	}

	//Note: Write header
	if err = utils.WriteHeader(fileoutput); err != nil {
		return err
	}

	if err = callMain(fileoutput, val); err != nil {
		return err
	}

	//Note: Compile all source files.
	for i := range(len(compiler.sourceFiles)) {
		if err := compiler.CompileSource(&compiler.sourceFiles[i], fileoutput); err != nil {
			return err
		}
	}

	log.Printf("Output file %s", compiler.OutputName)

	return nil
}

//Note: Every executable bytecode file should have "CALL main" as the first instruction.
func callMain(fileoutput *os.File, address int32) error {
	var opCode = isa.INST_CALL

	if err := binary.Write(fileoutput, binary.LittleEndian, byte(opCode)); err != nil {
		return err
	}
	
	if err := binary.Write(fileoutput, binary.LittleEndian, int32(address)); err != nil {
		return err
	}

	return nil
}

func (compiler *Compiler)CompileSource(file *SourceFile, fileoutput *os.File) error {
	log.Printf("Compiling File %s...\n", file.filename)

	var scanner *bufio.Scanner = bufio.NewScanner(strings.NewReader(file.content))
	var linecnt int32 = 0
	/*
	* Note: For every line, trim left-right(It is not needed because we trimmed the in the preprocessor, but i do what i want), we translate the
	* opcode and we write to a binary file the translated byte. If the opcode contains a value then we also write the int32 value in little endian format.
	* If the value is a label then we search the labels map and replace the label with the ip.
	*/
	
	for scanner.Scan() {
		linecnt++
		var line string = scanner.Text()

		line = strings.TrimLeft(line, " \t");
		line = strings.TrimRight(line, " \t");

		if line == "" || line[0] == '#' || line[0] == '!' || line[len(line) - 1] == ':' {
			continue
		}

		var tokens []string = strings.Fields(line)

		opCode, ok := isa.StringToInst[tokens[0]]

		if ok == false {
			if val, ok := compiler.macros[tokens[0]]; ok == true {
				opCode = isa.StringToInst[val]
			} else {
				var errstr string = fmt.Sprintf("%s:%d: Uknown Instruction %s", file.filename, linecnt, tokens[0]) 
				return errors.New(errstr)
			}
		}

		if err := binary.Write(fileoutput, binary.LittleEndian, byte(opCode)); err != nil {
			return err
		}

		if isa.InstHasOperand[opCode] == true {
			var val int32 = 0

			if len(tokens) < 2 {
				var errstr string = fmt.Sprintf("%s:%d: Opcode %s requires operand",  file.filename, linecnt, isa.InstToString[opCode]) 
				return errors.New(errstr)
			}

			var containsLabel = (opCode == isa.INST_JMP || opCode == isa.INST_JMPZ || opCode == isa.INST_CALL)

			if  containsLabel && (unicode.IsDigit(rune(tokens[1][0])) == false) {
				val, ok = compiler.labels[tokens[1]]

				if ok == false {
					if label, ok := compiler.macros[tokens[1]]; ok == true {
						val = compiler.labels[label]
					} else {
						var errstr string = fmt.Sprintf("%s:%d: Label %s does not exist",  file.filename, linecnt, tokens[1]) 
						log.Fatal(errstr)
					}
				}
			} else {
				var val64 int64 = 0
				if label, ok := compiler.macros[tokens[1]]; ok == true {
					val64, _ = strconv.ParseInt(label, 10, 32)
				} else {
					val64, _ = strconv.ParseInt(tokens[1], 10, 32)
				}

				val = int32(val64)
			}

			if err := binary.Write(fileoutput, binary.LittleEndian, int32(val)); err != nil {
				return err
			}
		}
	}

	log.Printf("File %v Compiled Successfully...\n", file.filename)

	return nil
}
