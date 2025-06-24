package main

import (
	"net/http"
	"fmt"
	"log"
)


func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from the Digit Occurrence API")
	})

	http.HandleFunc("/digit-occurrence", digitOccurrenceHandler)

	fmt.Println("Server listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
