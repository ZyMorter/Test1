package main

import (
	"log"
	"net/http"
	"time"
)


func greeting(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format("3:04 AM")  //get time
	currentDay := time.Now().Weekday().String() //get day of the week in string format 
	w.Write([]byte("Right now is "))
	w.Write([]byte(currentTime))
	w.Write([]byte("\nPlease enjoy the rest of your "))
	w.Write([]byte(currentDay))
}


func main() {
	// Create a new serveMux
	mux := http.NewServeMux()
	mux.HandleFunc("/greeting", greeting)
	log.Print("Starting server on :3000")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatal(err)

	}
} 

//http://localhost:4000/greeting
