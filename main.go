package main

import (
	"fmt"
	"io/ioutil"
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
		fmt.Fprint(w, "ALERT ALERT\n\n")

		url := "https://api.weather.gov/alerts/active/area/MO"

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprint(w, err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprint(w, err)
		}

		fmt.Fprint(w, string(body))
	})

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
