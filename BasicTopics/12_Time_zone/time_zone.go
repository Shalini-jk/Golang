package main 

import (
	"fmt"
	"time"
)

func main()  {
	current_time := time.Now().UTC()
	fmt.Println("current time :", current_time)

	//use of load loacation.

	get_location, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}

	get_time := current_time.In(get_location)
	//new_location := get_time.Location()
	//fmt.Println("new location: ", new_location)
	fmt.Println("time of new york with respect to america:",get_time)

	// Get the current time's time zone information
	currentLocation := time.Now().Location()
	fmt.Println("Current Time Zone:", currentLocation)

	// Get the time zone information of a specific time
	zoneName, zoneOffset := time.Now().Zone()
	fmt.Printf("Current Time Zone Name: %s, Offset: %d seconds\n", zoneName, zoneOffset)


}