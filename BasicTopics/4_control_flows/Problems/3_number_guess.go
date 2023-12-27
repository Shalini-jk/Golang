/*
Guess the Number:

Generate a random number between 1 and 100. Prompt the user to guess the number. Provide feedback to the user if their guess is too high or too low. Continue prompting until the user guesses the correct number.

*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1 and 100
	randomNumber := rand.Intn(100) + 1

	var input int
	var attempt int
	for {
		attempt ++
		fmt.Println("Enter your guess:")
		fmt.Scan(&input)

		if input == randomNumber {
			fmt.Println("Congratulations! You guessed correctly.")
			break
		} else if input < randomNumber {
			fmt.Println("Too low. Try again.")
		} else {
			fmt.Println("Too high. Try again.")
		}
	}
	fmt.Println("No of attempt :",attempt)
	fmt.Println("Random number :",randomNumber)
	fmt.Println("guessed number :",input)
	
}

