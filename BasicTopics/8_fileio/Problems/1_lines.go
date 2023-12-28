/*
problem:
Write a program that reads a text file and counts the number of lines in it. Use 
bufio.Reader for efficient line-by-line reading.

solution: 
we do not have any file to do these operations so first we create a file and write
conntent in it and then perform the aove operations.

Algorithm :
  1. create a file and then enrich the file with the content.
  2. then read the text from the created file and also count the number of lines.
  3. print the output that contain total numer of line. 
*/

package main

import (
    "fmt"
    "os"
	"bufio"
)

func main() {
    
	openfile, err := os.Open("sample.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer openfile.Close()

    // Create a buffered reader
    reader := bufio.NewReader(openfile)
	count := 0
	for {
		
        _, err := reader.ReadString('\n')
		count++
		
        if err != nil {
            break // End of file
			
        }
    }
	fmt.Println("Total number of lines",count)
	
	
}
