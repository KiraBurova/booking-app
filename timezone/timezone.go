package timezone

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var Timezones = []string{"Europe/Berlin", "Africa/Abidjan", "Africa/Addis_Ababa"}

func AddTimezone(w http.ResponseWriter, r *http.Request) {
	// read posted timezone
	timezone, _ := ioutil.ReadAll(r.Body)
	// set new timezone to timezones array
	Timezones = append(Timezones, string(timezone))
	json.NewEncoder(w).Encode(timezone)
}

func ListTimezones(w http.ResponseWriter, r *http.Request) {
	// get list of timezones in json format
	json.NewEncoder(w).Encode(Timezones)

}

func ConvertTimezone(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	timezone, _ := ioutil.ReadAll(r.Body)
	loc, _ := time.LoadLocation(string(timezone))
	fmt.Printf("Time in %v: %s\n", string(timezone), now.In(loc))
}
