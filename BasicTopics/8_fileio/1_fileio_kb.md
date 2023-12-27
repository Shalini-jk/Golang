File I/O (Input/Output) in Go is performed using the os and io packages from the Go Standard Library. These packages provide functions and types to interact with the file system, read from files, and write to files. Here's an overview of performing File I/O operations in Go:

# Reading from a File:
1. Opening a File:

Use the os.Open function to open a file for reading. It returns a file descriptor (*os.File) and an error.
```go

package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close() // Close the file when done
}

```
2. Reading File Content:

Use the Read or ReadAll functions from the io package to read content from the file.
```go

package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    content, err := ioutil.ReadFile("example.txt")
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }
    fmt.Println("File Content:", string(content))
}
```

# Writing to a File:

1. Creating or Opening a File for Writing:

Use the os.Create function to create a new file for writing or open an existing file. It returns a file descriptor.
```go

package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Create("output.txt")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close() // Close the file when done
}

```
2. Writing Content to a File:

Use the Write or WriteString functions from the io package to write content to the file.

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Create("output.txt")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()

    content := "Hello, File I/O in Go!"
    _, err = file.WriteString(content)
    if err != nil {
        fmt.Println("Error writing to file:", err)
        return
    }
    fmt.Println("Content written to file.")
}

```

# Checking File Existence:

1. Checking if a File Exists:

Use the os.Stat function to check if a file exists.
```go

package main

import (
    "fmt"
    "os"
)

func main() {
    _, err := os.Stat("example.txt")
    if os.IsNotExist(err) {
        fmt.Println("File does not exist.")
        return
    }
    fmt.Println("File exists.")
}
```

#Working with Directories:
1. Listing Files in a Directory:

Use the os.ReadDir function to list files in a directory.
```go

package main

import (
    "fmt"
    "os"
)

func main() {
    files, err := os.ReadDir(".")
    if err != nil {
        fmt.Println("Error reading directory:", err)
        return
    }
    for _, file := range files {
        fmt.Println(file.Name())
    }
}

```
These examples cover basic file I/O operations in Go. It's important to handle errors appropriately, and using the defer keyword ensures that files are closed even if an error occurs during processing. File I/O in Go is designed to be simple and efficient, providing a clean interface to interact with the file system.