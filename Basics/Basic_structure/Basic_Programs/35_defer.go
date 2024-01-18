// write a program to demonstrate defer keyword 
// defer keyword  is used to put the statement / data in stack and then execute it in lifo form(last in first out)

package main

import(
	"fmt"
)

func main() {
	defer fmt.Println("he shalini")
	fmt.Println("hey rajashri")
	defer fmt.Println("hey chitra")
}