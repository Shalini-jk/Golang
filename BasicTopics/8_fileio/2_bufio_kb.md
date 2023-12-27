In Go, buffered I/O is a technique used to efficiently read from or write to a file or stream by reducing the number of system calls and minimizing data transfers between the Go program and the operating system. The bufio package in the Go Standard Library provides a set of buffered I/O functions that wrap an io.Reader or io.Writer with buffering capabilities.

# Buffered Reading with bufio.Reader:
1. Creating a Buffered Reader:

Use the bufio.NewReader function to create a buffered reader. It takes an io.Reader as an argument.
```go

package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    // Create a buffered reader
    reader := bufio.NewReader(file)
}
```

2. Reading Line by Line:

Use the ReadString method to read a line (up to and including the delimiter) from the buffered reader.
```go

package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    // Create a buffered reader
    reader := bufio.NewReader(file)

    // Read and print lines
    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            break // End of file
        }
        fmt.Print(line)
    }
}
```

# Buffered Writing with bufio.Writer:

1. Creating a Buffered Writer:

Use the bufio.NewWriter function to create a buffered writer. It takes an io.Writer as an argument.
```go

package main

import (
    "bufio"
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

    // Create a buffered writer
    writer := bufio.NewWriter(file)
}
```
2. Writing with Buffer Flushing:

Use the Write method to write data to the buffered writer. The data is not immediately written to the underlying writer; it is stored in the buffer until it's full or until the Flush method is called.
```go

package main

import (
    "bufio"
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

    // Create a buffered writer
    writer := bufio.NewWriter(file)

    // Write data to the buffered writer
    data := "Hello, Buffered I/O in Go!"
    _, err = writer.WriteString(data)
    if err != nil {
        fmt.Println("Error writing to file:", err)
        return
    }

    // Flush the buffer to ensure data is written to the file
    err = writer.Flush()
    if err != nil {
        fmt.Println("Error flushing buffer:", err)
        return
    }
}
```

Buffered I/O is beneficial when dealing with large files or streams, as it helps minimize the number of system calls and improves performance. It is particularly useful in scenarios where reading or writing small chunks of data at a time is inefficient.

# reading from stdin using bufio
To read from standard input (stdin) using bufio in Go, you can create a bufio.Reader that wraps os.Stdin. This buffered reader allows you to efficiently read lines or other data from the standard input. Here's an example of reading lines from stdin:

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Create a buffered reader wrapping os.Stdin
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter text (press Ctrl+D to end input): \n")

	// Read lines from stdin until an error occurs (e.g., EOF or other input error)
	for {
		// ReadString reads until the first occurrence of delimiter '\n' or an error occurs
		line, err := reader.ReadString('\n')
		if err != nil {
			break // End of input or error
		}

		// Print the read line
		fmt.Print("Read: " + line)
	}
}
```

In this example:

We create a bufio.Reader wrapping os.Stdin.
We use a loop to continuously read lines from stdin until an error occurs (e.g., when the user presses Ctrl+D to signal the end of input).
The ReadString method is used to read a line until the newline character ('\n'). If you want to read individual runes or bytes, you can use other methods provided by bufio.Reader.
To run this program, you can input text line by line, and the program will print each line as it is read. Press Ctrl+D (or Ctrl+Z on Windows) to signal the end of input and exit the program.