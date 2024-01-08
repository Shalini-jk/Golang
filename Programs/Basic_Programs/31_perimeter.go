// write a program to find perimeter of a circle.
// perimeter : 2(pie)r

package main

import(
	"fmt"
)

func main() {
	//declaring variables
	var radius_circle float64
	var perimeter_circle float64

	//taking input from the user
	fmt.Println("Input the radius of circle")
	_,err := fmt.Scanln(&radius_circle)
	if err != nil {
		fmt.Println("Input is not in correct way, input the radius in float or integer",err)
	}

	//calculate the perimeter of the circle
	perimeter_circle = (2 * 3.14) * radius_circle
	fmt.Println("The perimeter of the circle :",perimeter_circle)
}
