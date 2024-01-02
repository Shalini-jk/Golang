package main
import (
	"fmt"
	"time"
)
func main() {
	// declaration of variable
	startTime := time.Now()
	var celsius float64
	var fahrenheit float64

	// taking input from the user
	fmt.Println("Enter the temperature in Celsius")
	fmt.Scan(&celsius)

	// Conversion of temperature from celsius to fahrenheit
	fahrenheit = (9.0/5.0) * celsius + 32

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	// print the temperature in celsius and fahrenhiet
	fmt.Printf("The Temperature in Celsius: %.2f °C\n", celsius)
    fmt.Printf("The Temperature in Fahrenheit: %.2f °F\n", fahrenheit)
	fmt.Println("elapsed time (Total time to complete the prgram):", elapsedTime)
	

}