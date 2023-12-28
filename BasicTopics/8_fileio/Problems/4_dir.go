/*
Write a program that generates a tree-like structure representing the hierarchy of files and directories starting from a specified root directory. Use os and os.FileInfo to gather information about files and directories.
*/

Package main

import (
    "fmt"
    "os"
	"io"
)

func main() {
    files, err := os.ReadDir("8_fileio")
    if err != nil {
        fmt.Println("Error reading directory:", err)
        return
    }
    for _, file := range files {
        fmt.Println(file.Name())
    }
}
