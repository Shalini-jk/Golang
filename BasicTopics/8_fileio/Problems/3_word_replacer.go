/*
problem :
Build a program that reads a text file, replaces occurrences of a specified word with 
another word, and then writes the modified content back to the file.

Algorithm :
  1. first we have to open the file
  2. we have to read the file line by line and check whether that specified word is present in that  line or not.
  3. If the specified word exist then we replaced that word with desired word.
*/


package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func replaceInFile(filePath, oldSubstring, newSubstring string) error {
	// Read the content of the file
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Perform the replacement using strings.Replace
	modifiedContent := strings.Replace(string(content), oldSubstring, newSubstring, -1)

	// Write the modified content back to the file
	err = ioutil.WriteFile(filePath, []byte(modifiedContent), 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Replacement in file '%s' completed.\n", filePath)
	return nil
}

func main() {
	filePath := "dest.txt"
	oldSubstring := "shalini"
	newSubstring := "malini"

	err := replaceInFile(filePath, oldSubstring, newSubstring)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
