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
	//sanitizedMAC := *userInput

	if strings.ContainsAny(*userInput,separators) {

		for _, r := range separators {
			*userInput = strings.ReplaceAll(*userInput, string(r), "")
		}
	}
	return strings.ToLower(*userInput)
}


func main() {
	mac := readInputMAC()
	
	if isValidMAC(&mac) {
		fmt.Printf("Great MAC address dude!\n")
		sanitizedmac := sanitizeInputMAC(&mac)
		fmt.Printf(sanitizedmac)
	}
}
