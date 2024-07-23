package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

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
	http.HandleFunc("/alerts", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "WEATHER ALERTS!!!\n\n")

		url := "https://api.weather.gov/alerts/active/area/MO"

		// Make a Get request to NWS
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprint(w, err)
		}

		// Read in response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprint(w, err)
		}

		// Unmarshal response body into struct
		var alerts Alerts
		err = json.Unmarshal(body, &alerts)
		if err != nil {
			fmt.Fprint(w, err)
		}

		// Print weather alerts and instructions to page
		for i := range alerts.Features {
			fmt.Fprintf(w, "Alert[%d]: %s\n", i, alerts.Features[i].Properties.Headline)
			fmt.Fprintf(w, "Instruction: %s\n\n", alerts.Features[i].Properties.Instruction)
		}

	})

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
