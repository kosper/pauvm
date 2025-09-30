package paudiss

import (
	"fmt"
	"os"
	"io"
	"errors"
	"encoding/binary"

	"github.com/kosper/pauvm/pkg/isa"
	"github.com/kosper/pauvm/pkg/utils"
)

func Dissassemble(flags *PDFlags) error {
	var file, err = os.Open(flags.filename)
	defer file.Close()

	if err != nil {
		return err
	}

	//Note: Read Header
	header, err := utils.ReadHeader(file)

	if err != nil {
		return err
	}

	//Note: Check that magic number is correct.
	if err = utils.CheckMagicNumber(header); err != nil { 
		return err; 
	}

	var instruction []byte = make([]byte, 1) 

	fmt.Printf("Pau Version: %v.%v.%v\nProgram:\n\n", header.Version[0], header.Version[1], header.Version[2])

	var ip int32 = 0
	
	for  {
		_, err := io.ReadFull(file, instruction)

		if err == io.EOF { return nil }
		if err != nil { return err }

		opcode, ok := isa.InstToString[isa.InstructionType(instruction[0])]

		if !ok {
			return errors.New(fmt.Sprintf("Can't dissassemble file, uknown instruction %d", instruction[0]))
		}

		var value int32 = 0

		if isa.InstHasOperand[isa.InstructionType(instruction[0])] {
			if err := binary.Read(file, binary.LittleEndian, &value); err != nil {
				return err
			}

			fmt.Printf("%5d | %v %v\n", ip, opcode, value)
		} else {
			fmt.Printf("%5d | %v\n", ip, opcode)
		}

		ip++
	}
}
