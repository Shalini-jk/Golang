package string

import (
	"fmt"
)

func Stringmanipulation() {
	fmt.Print("Enter a string: ")
	var userInput string
	fmt.Scanln(&userInput)

	// Reverse the string
	runes := []rune(userInput)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Print the reversed string
	reversedString := string(runes)
	fmt.Printf("Reversed string: %s\n", reversedString)
}

func Stringcount() {
	fmt.Print("Enter a string: ")
	var userInput string
	fmt.Scanln(&userInput)

	// Count the number of characters in the string
	characterCount := len(userInput)

	// Print the result
	fmt.Printf("Number of characters: %d\n", characterCount)


}