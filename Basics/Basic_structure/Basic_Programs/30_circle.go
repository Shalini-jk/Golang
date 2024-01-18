//write a program to calculate the area of a circle.
// area = radius*radius

package main

import(
	"fmt"
)

func main() {
	//declaring variables
	var radius_circle float64
	var area_circle float64

	//taking input from user
	fmt.Println("Enter the radius of circle")
	_, err := fmt.Scanln(&radius_circle)
	if err != nil {
		fmt.Println("incorrect input,radius must be integer or float",err)
	}

	//calculating the area of a circle
	area_circle = radius_circle * radius_circle
	fmt.Println("The radius of the circle is :", area_circle)
}