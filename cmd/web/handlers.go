package main

import (
	"encoding/json"
	"jakataGo/pkg/models"
	"log"
	"net/http"
)

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/public/images/favicon.ico")
}

func reviews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := dbConn()

	revs := []models.Review{}
	db.Find(&revs)

	db.Close()

	json, err := json.Marshal(revs)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}

func blogs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := dbConn()
	blogs := []models.BlogPost{}
	db.Find(&blogs)
	db.Close()

	json, err := json.Marshal(blogs)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}

func team(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := dbConn()
	team := []models.TeamMember{}
	db.Where("Salon = 1").Order("position asc").Find(&team)
	db.Close()

	json, err := json.Marshal(team)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}
