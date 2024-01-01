package main    

import (
	"calculator/operation"
	"fmt"
	
)

func main(){
	var choice int 
	for {
		fmt.Println("\n----Calculator Program----")
		fmt.Print("1.Addition\n2.Sutraction\n3.Multiplication\n4.Division\n")
		fmt.Println("Enter your choice : ")
		fmt.Scan(&choice)
  
	
	switch choice {
	case 1:
		fmt.Println(operation.Add(5,3))
	case 2:
		operation.Subtraction()
	case 3:
		operation.Multiplication()
	case 4:
		operation.Division()
	}
}
	
}