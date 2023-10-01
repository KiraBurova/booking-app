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
	timezone, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	Timezones = append(Timezones, string(timezone))
	json.NewEncoder(w).Encode(timezone)
}

func ListTimezones(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Timezones)

}

func ConvertTimezone(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	timezone, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	loc, loadLocationErr := time.LoadLocation(string(timezone))

	if loadLocationErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Printf("Time in %v: %s\n", string(timezone), now.In(loc))
}
