package isa

//Note: Cannot declare arrays as const.
//Note: Globals are not always bad. I don't care what you think.
var MagicNumber [4]byte = [4]byte{ '2', '0', '2', '5'}
var Version [3]byte = [3]byte{ 0, 0, 0 }

type FileHeader struct {
	MagicNumber [4]byte
	Version [3]byte
}
