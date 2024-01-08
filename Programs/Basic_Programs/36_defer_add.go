// writ a program to demonstrate defer in aspect to functions

package main 

import(
	"fmt"
)

//declare variable
const first_number = 34
const second_number = 78

func first_value () {
	fmt.Println(first_number)
}
func sum () {
	add := first_number + second_number
	fmt.Println("add of two number",add)
}
func main () {
	defer first_value()
	sum()
}