/* Implement a function that counts the number of characters (runes) in a UTF-8 encoded 
string. Use the utf8 package to handle Unicode characters correctly.
 */


 package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	
	utf8String := "Hello, 你好, こんにちは"
	runeCount := 0

	for len(utf8String) > 0 {
		
		_, size := utf8.DecodeRuneInString(utf8String)
		runeCount++
		utf8String = utf8String[size:]
	}

	// Display the result
	fmt.Printf("Number of runes: %d\n", runeCount)
}
