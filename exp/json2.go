package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Review struct {
	Review string
	Client string
	Stylist string
}

func main() {

	http.HandleFunc("/api/reviews", reviews)

	http.ListenAndServe(":8060", nil)
}

func reviews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	r1 := Review{
		"Great job!",
		"Anna Alexander",
		"Nat",
	}
	json, err := json.Marshal(r1)
	if err != nil {
		log.Println(err)
	}

	w.Write(json)
}
