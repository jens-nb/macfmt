package main

import (
	"fmt"
	"os"
	"strings"
)

// readInputMAC reads the input MAC address as command line argument.
func readInputMAC() string {
	userInput := os.Args[1]
	return userInput
}

func usage() {

	fmt.Println("Usage: macfmt <mac-address> <format>")
}

// isAllowedCharacter checks if a rune is one of the allowed characters
// in a MAC address.
func isAllowedCharacter(r *rune) bool {
	allowedCharacters := "abcdefABCDEF0123456789:.-"

	if !strings.ContainsRune(allowedCharacters, *r) {
		fmt.Println("Error: MAC address contains illegal character.")
		return false
	}
	return true
}

// isValidMAC checks if the provides string is a valid MAC address.
// Allowed Characters are numbers, digits, ".", ":", "-".
// Allowed positions of separators: 2,4,6,8,10, also
// Longest possible MAC has 17 chars: AB:CD:EF:78:90:12
// shortest possible MAC has 12 chars: ABCDEF012345
func isValidMAC(userInput *string) bool {

	for _, r := range *userInput {
		if !isAllowedCharacter(&r) {
			return false
		}
	}

	if len(*userInput) < 12 || len(*userInput) > 17 {
		fmt.Println("Error: MAC address has illegal length.")
		return false
	}
	return true
}

// sanitizeInputMAC removes all separators from an input MAC address
// and converts all letters to lowercase.
func sanitizeInputMAC(userInput *string) string {
	separators := ":.-"

	if strings.ContainsAny(*userInput, separators) {

		for _, r := range separators {
			*userInput = strings.ReplaceAll(*userInput, string(r), "")
		}
	}
	return strings.ToLower(*userInput)
}

// formatMAC takes a sanitized MAC as input and returns (prints) a MAC address in the specified format.
func formatMAC(sanitizedInput string, format string) {

	switch format {

	// Cisco format: abcd.ef01.2345
	case "cisco":
		var substrings []string
		s1 := sanitizedInput[0:4]
		s2 := sanitizedInput[4:8]
		s3 := sanitizedInput[8:12]
		substrings = append(substrings, s1,s2,s3) 

		fmt.Println(strings.Join(substrings, "."))
	}
}

func main() {
	mac := readInputMAC()

	if len(os.Args) < 3{
		usage()
		os.Exit(1)
	}

	if isValidMAC(&mac) {
		sanitizedmac := sanitizeInputMAC(&mac)
		formatMAC(sanitizedmac, os.Args[2])
	}
}
