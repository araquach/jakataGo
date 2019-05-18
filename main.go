package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"static-site/views"
)

var homeView *views.View
var teamView *views.View
var team_indView *views.View
var blogView *views.View
var blog_indView *views.View
var detailsView *views.View
var contactView *views.View

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

func blog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(blogView.Render(w, nil))
}

func blog_ind(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(blog_indView.Render(w, nil))
}

func details(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(detailsView.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))
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
	blogView = views.NewView("main", "views/blog.gohtml")
	blog_indView = views.NewView("main", "views/blog_ind.gohtml")
	detailsView = views.NewView("main", "views/details.gohtml")
	contactView = views.NewView("main", "views/contact.gohtml")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/team", team)
	r.HandleFunc("/team_ind", team_ind)
	r.HandleFunc("/blog", blog)
	r.HandleFunc("/blog_ind", blog_ind)
	r.HandleFunc("/details", details)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":3000", r)
}
