/*
Currency Converter:
Problem :
Create a program that converts an amount in one currency to another. Prompt the user to enter 
an amount in one currency (e.g., USD) and a conversion rate to another currency (e.g., EUR). 
Perform the conversion and display the result.Allow reverse conversion too.
formula : EUR=USD×Exchange Rate(exchange rate =8.5) USD=EUR×Exchange Rate(exchange rate = 1.20)
Algorithm :
   1. take a input from the user in one type of currency.
   2. then use the formula of currency to convert the currency from one type to another.
   3. print both the output at last. 
   4. then again reverse the conversion.
   5. print again both output.
*/
// lets convert usd to eur and eur to usd
package main
import (
	"fmt"
)
func main() {
	// variable declaration 
	var usd float64
	var eur float64

	fmt.Println("....currency converter from US Dollar to EUR and Vise Versa......")
	// Take input from user
	fmt.Println("Enter the currency in US Dollar")
	fmt.Scan(&usd)

	// conversion from usd to eur
	eur = usd * 0.85

	// print both the output
	fmt.Println("......US Dollar to EURO..........")
	fmt.Printf("Your Currency in US Dollar: $ %.2f \n", usd)
    fmt.Printf("Your: Currency in EUR : € %.2f \n", eur)

    // conversion from euro to usd
	usd = eur * 1.18

	// print oth the output
	fmt.Println("......EURO to US Dollar.........")
	fmt.Printf("Your: Currency in EUR : € %.2f \n", eur)
	fmt.Printf("Your Currency in US Dollar: $ %.2f \n", usd)
}