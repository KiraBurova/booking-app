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

func addTimezone()     {}
func listTimezones()   {}
func convertTimezone() {}
