package main 
import (
	"fmt"
	"string/string"
	"os"
)

func main(){
	var choice int 
	for {
		fmt.Println("\n----String Manipulation Program---")
		fmt.Print("1.reverse the string\n2.count the no of character in string\n3.exit\n")
		fmt.Println("Enter your choice : ")
		fmt.Scan(&choice)
  
	
	switch choice {
	case 1:
		string.Stringmanipulation()
	case 2:
		string.Stringcount()
	case 3:
		os.Exit(1)
	
	}
}
	
}