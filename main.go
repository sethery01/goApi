package main

import (
	"fmt"
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

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
