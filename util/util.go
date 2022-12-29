package util

import (
	"errors"
)

// Split a string into equal sized chunks specified by size.
// Cisco format MAC addresses need 3x4 chunks, others 6x2.
func Chunk(input string, size int) ([]string, error) {

	if len(input)%size != 0 {
		return nil, errors.New("Error: input not divisible by chunk size.")
	}

	var result []string
	for i := 0; i < len(input); i += size {
		chunk := input[i : i+size]
		result = append(result, chunk)

	}

	return result, nil

}
