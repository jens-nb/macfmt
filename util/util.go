package util

import (
	"errors"
)

// Split a string into equal sized chunks specified by size.
// Cisco format MAC addresses need 3x4 chunks, others 6x2.
func Chunk(s string, size int) ([]string, error) {

	if len(s)%size != 0 {
		return nil, errors.New("Chunk error: s not divisible by chunk size.")
	}

	var result []string
	for i := 0; i < len(s); i += size {
		chunk := s[i : i+size]
		result = append(result, chunk)

	}

	return result, nil

}
