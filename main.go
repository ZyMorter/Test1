package main

import (
	"log"
	"net/http"
)

// The handler function:
func bio(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("\nMy name is Zyon Jamaal Morter, I'm 20 years old. I was born on 31/12/02 and am pursuing an associate degree in Information Technology at UB."))
}	

func main() {
	// Create a new serveMux
	mux := http.NewServeMux()
	mux.HandleFunc("/", bio)
	log.Print("Starting server on :2000")
	err := http.ListenAndServe(":2000", mux)
	if err != nil {
		log.Fatal(err)
	}
} 





