package isa

type Syscall byte

//Note: It will not be used by the user, only for debugging.
const (
	SYSCALL_EXIT Syscall = iota
)
