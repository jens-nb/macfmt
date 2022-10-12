package main

import (
	"fmt"
	"os"
)

func getAllowedSeparators()  []string{
	return []string{":", ".", "-"} }

func readInputMAC() string {
	userInput := os.Args[1]
	return userInput
}

func isValidMAC(userInput string) bool {
	var isValid bool

	return isValid
}

func main() {
	mac := readInputMAC()

	fmt.Printf("MAC address is %s", mac)
}
