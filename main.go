// main.go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type GPS struct {
	Lat interface{} `json:"lat"`
	Lng interface{} `json:"lng"`
}

var currentLocation GPS

func main() {
	mux := http.NewServeMux()
	log.Print("Create Server...")

	mux.HandleFunc("/", BaseHandler)
	mux.HandleFunc("/gps-data", GpsHandler)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err.Error())
	}
}
func BaseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not Implemented")
	return
}

func GpsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, currentLocation)
	}
	if r.Method == "POST" {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading request body: %v", err)
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}
		//Only work with application/json content type
		err = json.Unmarshal(body, &currentLocation)
		if err != nil {
			log.Printf("Failed to unmarshal response body: %v", err)
		}

		// // Respond with a success message
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Request payload received"))

		fmt.Println("Latitude:", currentLocation.Lat)
		fmt.Println("Longitude:", currentLocation.Lng)
		return
	}
	return
}
