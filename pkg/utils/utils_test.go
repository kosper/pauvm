package utils

import (
	"testing"
	"os"

	. "github.com/kosper/pauvm/pkg/isa"
)

func TestIsFileExtension(t *testing.T) {
	//Note: Basic test.
	result := IsFileExtension("test.pau", ".pau")

	if result != nil {
		t.Error(result)
	}

	//Note: Test that we can have multiple points in a filename.
	result = IsFileExtension("test.pau.pau", ".pau")

	if result != nil {
		t.Error(result)
	}

	//Note: Test that we can have extract file extension from a path.
	result = IsFileExtension("examples/path/test.pv", ".pv")

	if result != nil {
		t.Error(result)
	}

	//Note: Test that we cannot have an extension without a filename.
	result = IsFileExtension(".pv", ".pv")

	if result == nil {
		t.Error("Error: result should have been an error instead of nil!")
	}

	//Note: Test that we can have spaces in a filename.
	result = IsFileExtension("test .pv", ".pv")

	if result != nil {
		t.Error(result)
	}
}

//TODO: Move this inside isa/header_test.go
func TestMagicNumber(t *testing.T) {
	//Note: Evaluate that magic number is checked correctly.
	var header FileHeader

	header.MagicNumber[0] = '2'
	header.MagicNumber[1] = '0'
	header.MagicNumber[2] = '2'
	header.MagicNumber[3] = '5'

	result := CheckMagicNumber(&header)

	if result != nil {
		t.Error(result)
	}
}

func TestReadHeader(t *testing.T) {
	var header *FileHeader

	//Note: Evaluate that header is correct from examples/counter/counter.pau
	exampleFile, err := os.Open("../../examples/counter/counter.pau")

	if err != nil {
		t.Error(err)
		return
	}

	header, result := ReadHeader(exampleFile)

	if result = CheckMagicNumber(header); result != nil {
		t.Error(result)
	}

	if header.Version[0] != Version[0] && header.Version[1] != Version[1] && header.Version[2] != Version[2] {
		t.Errorf("Error, ReadHeader should've read %v, instead of %v\n", string(Version[:]), string(header.Version[:]))
	}

	exampleFile.Close()

	//Note: Evaluate that header is correct from examples/factorial/factorial.pau
	exampleFile, err = os.Open("../../examples/factorial/factorial.pau")

	if err != nil {
		t.Error(err)
		return
	}

	header, result = ReadHeader(exampleFile)

	if result = CheckMagicNumber(header); result != nil {
		t.Error(result)
	}

	if header.Version[0] != Version[0] && header.Version[1] != Version[1] && header.Version[2] != Version[2] {
		t.Errorf("Error, ReadHeader should've read %v, instead of %v\n", string(Version[:]), string(header.Version[:]))
	}

	exampleFile.Close()

	//Note: Evaluate that header is correct from examples/function_example/function.pau
	exampleFile, err = os.Open("../../examples/function_example/function.pau")

	if err != nil {
		t.Error(err)
		return
	}

	header, result = ReadHeader(exampleFile)

	if result = CheckMagicNumber(header); result != nil {
		t.Error(result)
	}

	if header.Version[0] != Version[0] && header.Version[1] != Version[1] && header.Version[2] != Version[2] {
		t.Errorf("Error, ReadHeader should've read %v, instead of %v\n", string(Version[:]), string(header.Version[:]))
	}

	exampleFile.Close()

	//Note: Evaluate that header is correct from examples/locals_example/locals.pau
	exampleFile, err = os.Open("../../examples/locals_example/locals.pau")

	if err != nil {
		t.Error(err)
		return
	}

	header, result = ReadHeader(exampleFile)

	if result = CheckMagicNumber(header); result != nil {
		t.Error(result)
	}

	if header.Version[0] != Version[0] && header.Version[1] != Version[1] && header.Version[2] != Version[2] {
		t.Errorf("Error, ReadHeader should've read %v, instead of %v\n", string(Version[:]), string(header.Version[:]))
	}

	exampleFile.Close()
}
