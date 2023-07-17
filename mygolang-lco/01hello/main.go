package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/success", successHandler)
	http.HandleFunc("/unauthorized", unauthorizedHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func successHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Success!",
	}

	// Set the appropriate headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Convert the response struct to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the JSON response
	w.Write(jsonResponse)
}

func unauthorizedHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Unauthorized!",
	}

	// Set the appropriate headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)

	// Convert the response struct to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the JSON response
	w.Write(jsonResponse)
}
