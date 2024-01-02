package main 

import (
	"fmt"
	"time"
)

func main()  {
	//Get the current time

	currentTime := time.Now()
	fmt.Println("Current Time:", currentTime)
	fmt.Printf("Type of curent time :%T \n",currentTime)

	// Format time as a string
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	fmt.Println("Formatted Time:", formattedTime)
	fmt.Printf("Type of formated time :%T \n",formattedTime) //this is used to print the type of time
    
	// using parse function
	parsedTime, err := time.Parse("2006-01-02 15:04:05", formattedTime)
	fmt.Println("parsed Time:", parsedTime)
	fmt.Printf("Type of parsed time :%T \n",parsedTime) 
	if err != nil {
		fmt.Println("Error parsing time:", err)
	}

	// getting current day
	currentTime1 := time.Now()
	current_day := currentTime1.Weekday()
	fmt.Println("current day :", current_day)
	weekname := current_day.String()[:3]
	fmt.Println("current day  name:", weekname)

	//getting current month

	currentMonth := currentTime.Month()
	fmt.Println("Current Month:", currentMonth)

	//getting current year

	currentYear := currentTime.Year()
	fmt.Println("Current Year:", currentYear)

	// Add a duration
	newtime := currentTime.Add(time.Hour * 22)
	fmt.Println("Future Time (after 24 hours):", newtime)







}