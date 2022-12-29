package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jens-nb/macfmt/util"
)

// readInputMAC reads the input MAC address as command line argument.
func readInputMAC() string {
	userInput := os.Args[1]
	return userInput
}

// usage prints a help message in case of input error.
func usage() {

	fmt.Println("Usage: macfmt <mac-address> <format>")
}

// isAllowedCharacter checks if a rune is one of the allowed characters
// in a MAC address.
func isAllowedCharacter(r rune) bool {
	allowedCharacters := "abcdefABCDEF0123456789:.-"

	if !strings.ContainsRune(allowedCharacters, r) {
		fmt.Println("Error: MAC address contains illegal character.")
		return false
	}
	return true
}

// isValidMAC checks if the provides string is a valid MAC address.
// Allowed Characters are numbers, digits, ".", ":", "-".
// Allowed positions of separators: 2,4,6,8,10
// Longest possible MAC has 17 chars: AB:CD:EF:78:90:12
// shortest possible MAC has 12 chars: ABCDEF012345
func isValidMAC(userInput string) bool {

	for _, r := range userInput {
		if !isAllowedCharacter(r) {
			fmt.Println("Invalid MAC address: contains invalid character.")
			return false
		}
	}

	if len(userInput) < 12 || len(userInput) > 17 {
		fmt.Println("Invalid MAC address: has invalid length.")
		return false
	}
	return true
}

// sanitizeInputMAC removes all separators from an input MAC address
// and converts all letters to lowercase.
func sanitizeInputMAC(userInput string) string {
	separators := ":.-"

	if strings.ContainsAny(userInput, separators) {

		for _, r := range separators {
			userInput = strings.ReplaceAll(userInput, string(r), "")
		}
	}
	return strings.ToLower(userInput)
}

// formatMAC takes a sanitized MAC as input and returns (prints) a MAC address in the specified format.
func formatMAC(sanitizedInput string, format string) (string, error) {

	var result string

	substr, err := util.Chunk(sanitizedInput, 2)
	if err != nil {
		return "", err
	}

	switch format {

	case ":":
		result = strings.Join(substr, ":")

	case "-":
		result = strings.Join(substr, "-")

	case "cisco":
		substr, err := util.Chunk(sanitizedInput, 4)
		if err != nil {
			return "", err
		}
		result = strings.Join(substr, ".")
	}
	return result, nil

}

func main() {

	if len(os.Args) < 3 {
		usage()
		os.Exit(1)
	}

	mac := readInputMAC()

	if isValidMAC(mac) {
		sanitizedMac := sanitizeInputMAC(mac)
		result, err := formatMAC(sanitizedMac, os.Args[2])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
	}
}
