package network

type Code int

// Define return code.
const (
	FAILURE Code = -1   // default failure code
	SUCCESS Code = iota // default success code
)
