package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jens-nb/macfmt/util"
)

// usage prints a help message in case of input error.
func usage() {
	fmt.Println("Usage: macfmt <MAC-address> <format>")
}

// isValid checks if the provides string is a valid MAC address.
// Allowed Characters are numbers, digits, ".", ":", "-".
// Longest possible MAC has 17 chars: AB:CD:EF:78:90:12
// shortest possible MAC has 12 chars: ABCDEF012345
func isValid(userInput string) bool {
	allowedCharacters := "abcdefABCDEF0123456789:.-"

	if len(userInput) < 12 || len(userInput) > 17 {
		fmt.Println("isValid: Invalid MAC address: has invalid length.")
		return false
	}

	for _, r := range userInput {
		if !strings.ContainsRune(allowedCharacters, r) {
			fmt.Println("isValid: Invalid MAC address: contains invalid character.")
			return false
		}
	}

	return true
}

// sanitize removes all separators from an input MAC address
// and converts all letters to lowercase.
func sanitize(macAddr string) string {
	separators := ":.-"

	if strings.ContainsAny(macAddr, separators) {
		for _, r := range separators {
			macAddr = strings.ReplaceAll(macAddr, string(r), "")
		}
	}
	return strings.ToLower(macAddr)
}

// format takes a sanitized MAC as input and returns (prints) a MAC address in the specified format.
func format(macAddr string, format string) (string, error) {
	var result string

	chunks, err := util.Chunk(macAddr, 2)
	if err != nil {
		return "", fmt.Errorf("format: failed to format MAC address: %w", err)
	}

	switch format {
	case ":":
		result = strings.Join(chunks, ":")
	case "-":
		result = strings.Join(chunks, "-")
	case "cisco":
		chunks, err := util.Chunk(macAddr, 4)
		if err != nil {
			return "", fmt.Errorf("format: failed to format MAC address: %w", err)
		}
		result = strings.Join(chunks, ".")
	}
	return result, nil
}

func main() {

	if len(os.Args) < 3 {
		usage()
		os.Exit(1)
	}

	mac := os.Args[1]

	if isValid(mac) {
		sanitizedMac := sanitize(mac)
		result, err := format(sanitizedMac, os.Args[2])

		if err != nil {
			fmt.Printf("Failed to format MAC: %s", err)
		}
		fmt.Println(result)
	}
}
