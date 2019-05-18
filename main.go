package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"static-site/views"
)

var homeView *views.View
var teamView *views.View
var team_indView *views.View

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

func team(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(teamView.Render(w, nil))
}

func team_ind(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(team_indView.Render(w, nil))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	homeView = views.NewView("main", "views/home.gohtml")
	teamView = views.NewView("main", "views/team.gohtml")
	team_indView = views.NewView("main", "views/team_ind.gohtml")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/team", team)
	r.HandleFunc("/team_ind", team_ind)
	http.ListenAndServe(":3000", r)
}
