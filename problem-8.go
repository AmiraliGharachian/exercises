package main

import (
	"fmt"
	"time"
)

func main() {
	var inputTime string

	fmt.Print("Enter the time (hh:mm:ss AM/PM): ")
	fmt.Scanln(&inputTime)

	layout := "03:04:05 PM"

	userTime, err := time.Parse(layout, inputTime)
	if err != nil {
		fmt.Println("Error: Invalid time format. Please use hh:mm:ss AM/PM format.")
		return
	}

	startEvening, _ := time.Parse(layout, "07:00:00 PM")
	endEvening, _ := time.Parse(layout, "09:00:00 PM")
	startMorning, _ := time.Parse(layout, "08:00:00 AM")
	endMorning, _ := time.Parse(layout, "10:00:00 AM")
	startMidnight, _ := time.Parse(layout, "12:00:00 PM")
	endMidnight, _ := time.Parse(layout, "06:00:00 AM")

	if (userTime.After(startEvening) && userTime.Before(endEvening)) ||
		(userTime.After(startMorning) && userTime.Before(endMorning)) ||
		(userTime.Equal(startMidnight) || userTime.Before(endMidnight) || userTime.After(startMidnight)) {
		fmt.Println("It's gaming time!")
	} else {
		fmt.Println("It's not the time for games yet...")
	}
}
