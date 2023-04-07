package main

import (
	"log"
	"math/rand"
	"net/http"
	
)


// store random quotes:
var quotes = []string{
	"The greatest glory in living lies not in never falling, but in rising every time we fall. -Nelson Mandela",
"The way to get started is to quit talking and begin doing. -Walt Disney",
"Your time is limited, so don't waste it living someone else's life. Don't be trapped by dogma  which is living with the results of other people's thinking. -Steve Jobs",
"If life were predictable it would cease to be life, and be without flavor. -Eleanor Roosevelt",
"If you look at what you have in life, you'll always have more. If you look at what you don't have in life, you'll never have enough. -Oprah Winfrey",
"If you set your goals ridiculously high and it's a failure, you will fail above everyone else's success. -James Cameron",
"Life is what happens when you're busy making other plans. -John Lennon",
"Two things are infinite: the universe and human stupidity; and I'm not sure about the universe. ― Albert Einstein",
"In three words I can sum up everything I've learned about life: it goes on.― Robert Frost",

}

func random(w http.ResponseWriter, r *http.Request) {
	// Get a random quote
	quote := quotes[rand.Intn(len(quotes))]
	//write the response
	w.Write([]byte(quote))
}
func main() {
	// Create a new serveMux
	mux := http.NewServeMux()
	mux.HandleFunc("/random", random)
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
} 

//http://localhost:4000/random