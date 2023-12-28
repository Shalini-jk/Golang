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
	"os"
	"io"
	"bufio"
)

func main() {
	sourcefile, err := os.Open("dest.txt")
	if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer sourcefile.Close()

	//Create a buffered reader
    reader := bufio.NewReader(sourcefile)
	for {
        _, err := reader.ReadString('\n')
		replacefile := string.replaced(sourcefile, "man", "animal")
        if err != nil {
            break // End of file
			
        }
    }
	fmt.Println(dest.txt)
}

