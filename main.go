package main

import (
	"fmt"
	"time"
)

func main() {
	timezone := getCurrentTimezone()
	fmt.Println(timezone)
}

func getCurrentTimezone() string {
	t := time.Now()
	zone, _ := t.Zone()

	return zone
}

func addTimezone() {
	// choose a timezone to add somehow ?? from a static list from somewhere, for example
	// add to array -> save somewhere -> database ??
}
func listTimezones() {
	// get list of timezones from database
}
func convertTimezone() {
	// get timezone getCurrentTimezone()
	// convert to any timezone from the listTimezones()
}
