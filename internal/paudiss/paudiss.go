package paudiss

import (
	"os"
	"fmt"
	"io"
	"errors"
	"encoding/binary"

	"github.com/kosper/pauvm/pkg/isa"
)

func Dissassemble(filename string) error {
	var file, err = os.Open(filename)
	defer file.Close()

	if err != nil {
		return err
	}

	var instruction []byte = make([]byte, 1) 

	fmt.Printf("Program:\n\n")
	for {
		_, err := io.ReadFull(file, instruction)

		if err == io.EOF { break }
		if err != nil { return err }

		opcode, ok := isa.InstToString[isa.InstructionType(instruction[0])]

		if !ok {
			return errors.New(fmt.Sprintf("Cant dissassembly file, uknown instruction %d", instruction[0]))
		}

		var value int32 = 0

		if isa.InstHasOperand[isa.InstructionType(instruction[0])] {
			if err := binary.Read(file, binary.LittleEndian, &value); err != nil {
				return err
			}

			fmt.Printf("%v %v\n", opcode, value)
		} else {
			fmt.Printf("%v\n", opcode)
		}
	}

	return nil
}
