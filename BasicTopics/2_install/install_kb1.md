Here are instructions to install the golang
1. install golang using apt
> sudo apt install golang-go
2. Check fot go getting installed
> go version
Your go version should be go1.18.1 or higher
3. Write your first hello.go program
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}

4. Run the program
> go run hello.go

5. Compile the go program
> go build hello.go
It will generate hello.
> ./hello
To run the binary

6. Check go env variables
> go env

