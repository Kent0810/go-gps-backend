// main.go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Gps struct {
    Lat interface{} `json:"lat"`
    Lng interface{} `json:"lng"`
}
var currentLocation Gps

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Print('?')
    })
    mux.HandleFunc("/gps-data", GpsHandler)
    err := http.ListenAndServe(":8080", mux)
    fmt.Print(err)
}
func BaseHandler(w http.ResponseWriter, r * http.Request) {
    fmt.Fprintf(w,"Not Implemented")
    return
}

func GpsHandler (w http.ResponseWriter, r *http.Request) {
    body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &currentLocation)
    if err != nil {
		log.Fatalf("Failed to unmarshal response body: %v", err)
	}

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Request payload received"))

	fmt.Println("Latitude:", currentLocation.Lat)
	fmt.Println("Longitude:", currentLocation.Lng)
    return
}
