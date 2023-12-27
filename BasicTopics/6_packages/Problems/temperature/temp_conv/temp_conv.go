package temp_conv

import (
	"fmt"
	
)

func Temperature_Conversion() {
	// declaration of variable
	var celsius float64
	var fahrenheit float64

	// taking input from the user
	fmt.Println("Enter the temperature in Celsius")
	fmt.Scan(&celsius)

	// Conversion of temperature from celsius to fahrenheit
	fahrenheit = (9.0/5.0) * celsius + 32

	// print the temperature in celsius and fahrenhiet
	fmt.Printf("The Temperature in Celsius: %.2f °C\n", celsius)
    fmt.Printf("The Temperature in Fahrenheit: %.2f °F\n", fahrenheit)
}


