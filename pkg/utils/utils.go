package utils 

import (
	"fmt"
	"os"
	"errors"
	"encoding/binary"

	"github.com/kosper/pauvm/pkg/isa"
)

//Note: Extracts file extension of filename and compares it with the desired extension.
func IsFileExtension(filename string, extension string) error {
	extensionIndex := 0

	//Note: Search from the end until a . is found, no need to check 0'th index. 
	for i := len(filename) - 1; i > 0; i-- {
		if filename[i] == '.' {
			extensionIndex = i;
			break;
		}
	}

	//Note: This means that stop point was never found, so file does not have an extension.
	if extensionIndex == 0 {
		var ferror string = fmt.Sprintf("File %s does not have extension.", filename)
		return errors.New(ferror)
	}

	var fileExtension string = filename[extensionIndex:]

	if  fileExtension == extension {
		return nil 
	}

	var ferror string = fmt.Sprintf("File %s should be of extension %s, instead is of extension %s", filename, extension, fileExtension);

	return errors.New(ferror)
}

func WriteHeader(fileoutput *os.File) error {
	//Note: I do this so i can write header in one call instead of two.
	var header [7]byte = [7]byte{ 
		isa.MagicNumber[0],
		isa.MagicNumber[1],
		isa.MagicNumber[2],
		isa.MagicNumber[3],

		isa.Version[0],
		isa.Version[1],
		isa.Version[2],
	}

	if err := binary.Write(fileoutput, binary.LittleEndian, [7]byte(header)); err != nil {
		return err
	}

	return nil
}

//Note: Extracts header of .pau bytecode file.
func ReadHeader(fileinput *os.File) (*isa.FileHeader, error) {
	var header isa.FileHeader

	if err := binary.Read(fileinput, binary.LittleEndian, &header); err != nil {
		return &header, err
	}

	return &header, nil
}

//Note: Compares headers magic number with the actual magic number.
func CheckMagicNumber(header *isa.FileHeader) error {
	if string(isa.MagicNumber[:]) == string(header.MagicNumber[:]) {
		return nil
	}

	var ferror string = fmt.Sprintf("Magic number is wrong!, Should be %s, is %s", string(isa.MagicNumber[:]), string(header.MagicNumber[:]))

	return errors.New(ferror)
}
