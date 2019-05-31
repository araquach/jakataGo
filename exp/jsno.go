package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Person struct {
	Fname string
	Lname string
	Things []string
}


func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/mshl", mshl)
	http.HandleFunc("/encd", encd)
	http.ListenAndServe(":3010", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	s := `<!DOCTYPE HTML>
			<html lang = "en">
			<head>
			  <title>Foo</title>
			  <meta charset = "UTF-8" />
			</head>
			<body>
			  <h1>You are at Foo</h1>
			</body>
			</html>`
	w.Write([]byte(s))
}

func mshl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := Person{
		"Adam",
		"Carter",
		[]string{"Potatoe", "Car", "Wood"},
	}
	json, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)


}

func encd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := Person{
		"Adam",
		"Carter",
		[]string{"Potatoe", "Car", "Wood"},
	}
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println(err)
	}
}

