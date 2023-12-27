/*
Rock, Paper, Scissors Game:

Implement a simple text-based Rock, Paper, Scissors game. Prompt the user to enter their 
choice (Rock, Paper, or Scissors) and generate a random choice for the computer. 
Determine the winner based on the rules: Rock crushes Scissors, Scissors cuts Paper,
Paper covers Rock. Both having same choice is a tie.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// get user choice
	var userchoice string
	fmt.Println("Enter choice")
	fmt.Scan(&userchoice)

	// get random choice 
	rand.Seed(time.Now().UnixNano())
	choices := []string{"Rock", "Paper", "Scissors"}
	randomChoice := choices[rand.Intn(len(choices))]

	fmt.Printf("You chose %s.\n", userchoice)
	fmt.Printf("Computer chose %s.\n", randomChoice)

// determine the winners
	if (randomChoice == userchoice){
		fmt.Println("Its a tie !!")
	}else if(userchoice == "Rock" && randomChoice == "Scissors") {
		fmt.Println("You win!")
	}else if(userchoice == "Scissors" && randomChoice == "Paper"){
		fmt.Println("you win !!")
	}else if(userchoice == "Paper" && randomChoice == "Rock"){
		fmt.Println("you win !!")
	}else{
		fmt.Println("computer win !!")
	}
	

}
