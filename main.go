package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Define a global array to check for state codes
var stateCodes = []string{
	"AL", "AK", "AZ", "AR", "CA", "CO", "CT", "DE", "FL", "GA", "HI", "ID", "IL", "IN", "IA",
	"KS", "KY", "LA", "ME", "MD", "MA", "MI", "MN", "MS", "MO", "MT", "NE", "NV", "NH", "NJ",
	"NM", "NY", "NC", "ND", "OH", "OK", "OR", "PA", "RI", "SC", "SD", "TN", "TX", "UT", "VT",
	"VA", "WA", "WV", "WI", "WY",
}

// Function to get weather alerts by state from the NWS
func getAlerts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "WEATHER ALERTS!!!\n\n")

	state := r.URL.Query().Get("state")
	for i := 0; i <= 50; i++ {
		if i == 50 {
			fmt.Fprint(w, "ERROR: Incorrect state code or no state code given.\n\n")
			return
		} else if state == stateCodes[i] {
			break
		} else {
			continue
		}
	}

	url := fmt.Sprintf("https://api.weather.gov/alerts/active/area/%s", state)

	// Make a Get request to NWS
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	// Read in response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	// Unmarshal response body into struct
	var alerts Alerts
	err = json.Unmarshal(body, &alerts)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	// // Encode books data as JSON
	// jsonData, err := json.Marshal(alerts)
	// if err != nil {
	// 	fmt.Fprint(w, err)
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.Write(jsonData)

	// Print weather alerts and instructions to page
	for i := range alerts.Features {
		fmt.Fprintf(w, "Alert[%d]: %s\n", i, alerts.Features[i].Properties.Headline)
		fmt.Fprintf(w, "Instruction: %s\n\n", alerts.Features[i].Properties.Instruction)
	}
}

func main() {

	// Kubernetes health route.
	http.HandleFunc("/livez", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Healthy\n")
	})

	// Kubernetes ready route.
	http.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Ready\n")
	})

	// Get weather alerts
	http.HandleFunc("/alerts", getAlerts)

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
