package random

import (
	"fmt"
	"time"
	"math/rand"
)

func Random_number_generator() {
	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1 and 9999
	randomNumber := rand.Intn(9999) + 1
	fmt.Println("The random number", randomNumber)
}


