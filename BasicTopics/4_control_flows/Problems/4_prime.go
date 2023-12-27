/*
Prime number check: Check if given number is prime or not
*/

/*
Describe problem: prime number is a number which is onl divisile  1 and the numer itself
Algorithm : 
       1. first we take a input from the user
	   2. Then we can check it is completely divisible by 1 and itself only and not by other number
	   3. if it full fill the condition then it will e prme number other wise not a prime number
*/
// program

package main

import(
	"fmt"
	"math"
)

func main(){
	// variable declaration
	var input int
	var prime = true

	// take user input
	fmt.Println("Enter the number which you want to check :")
	fmt.Scan(&input)

	if (input < 1){
		prime = false
		fmt.Println(" 1 is not a prime number")
	}else{

	// get square root of a number using math package
	squinput := math.Sqrt(float64(input))
	fmt.Println("Square root of the given number",squinput)
	issqu := int(squinput)
	fmt.Println("Square root of the given number",issqu)

	// write the logic to check the number is prime number or not
	for i := 2 ; i <= issqu ; i++ {
		if input % i == 0 {
			prime = false
			break
			
		}
	}
}
if prime == false {
	fmt.Println(" The number is not a prime number")

}else{
	fmt.Println(" The number is a prime number")
}
	
}



