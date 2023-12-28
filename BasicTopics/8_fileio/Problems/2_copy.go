/*
Create a program that copies the content of one file to another. Allow the user to specify 
the source and destination file paths.
*/

package main 
import (
	"fmt"
	"os"
	"io"
)

func main() {
	// opening the source file 
	var file1, file2 string 

	fmt.Println("Enter the name of source file")
	fmt.Scan(&file1)

	fmt.Println("Enter the name of destination file")
	fmt.Scan(&file2)

	sourcefile, err := os.Open(file1)
	if err != nil {
        fmt.Println("This file does not eist:", err)
        return
    }
    defer sourcefile.Close()

	destinationfile, err := os.Create(file2)
	if err != nil {
		fmt.Println("This file not exist", err)
		return
	}
	defer destinationfile.Close()

	io.Copy(destinationfile, sourcefile)
	fmt.Println("File copied  sucessfull from source file to destination file")
	
}

