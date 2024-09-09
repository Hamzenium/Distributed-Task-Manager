package p2p

import (
	"encoding/gob" // Imports the gob package for binary encoding/decoding
	"io"           // Imports the io package for input/output operations
)

// Decoder interface defines a contract for decoding data from an io.Reader
// into a provided variable of any type, returning an error if the decoding fails.
type Decoder interface {
	Decode(io.Reader, any) error
}

// GOBDecoder is a struct that implements the Decoder interface using
// Go's gob package for binary decoding.
type GOBDecoder struct{}

// Decode method on GOBDecoder takes an io.Reader and a variable of any type.
// It uses gob.NewDecoder to decode the data from the reader into the variable.
// Returns an error if the decoding process fails.
func (dec GOBDecoder) Decode(r io.Reader, v any) error {
	// Creates a new gob.Decoder with the provided io.Reader and decodes
	// the data into the variable v.
	return gob.NewDecoder(r).Decode(v)
}