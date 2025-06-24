//server.go - Entry point for the API server.

package main

import (
	"net/http"
	"fmt"
	"log"
)

//main initialization for the HTTP server, sets up routes.
func main() {
	//Inits our Root endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Digit Occurrence API")
	})
	//adds primary endpoint, logic found in src/digitOccurrence
	http.HandleFunc("/digit-occurrence", digitOccurrenceHandler)

	fmt.Println("Server listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
