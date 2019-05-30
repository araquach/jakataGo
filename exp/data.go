package main

import (
	"adsTest/views"
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

var homeView *views.View
var contactView *views.View
var db *sql.DB

type Blog struct {
	gorm.Model
	Title string
	Slug string
	Author string
	Para string
}

func home(w http.ResponseWriter, r *http.Request) {
	b1 := Blog{
		Title : "Blog Post 1",
		Slug : "blog-post-one",
		Author : "Nat",
		Para : "This is the first para of b1",
	}
	b2 := Blog{
		Title : "Blog Post 2",
		Slug : "blog-post-two",
		Author : "Nat",
		Para : "This is the first para of b2",
	}
	b3 := Blog{
		Title : "Blog Post 3",
		Slug : "blog-post-two",
		Author : "Nat",
		Para : "This is the first para of b3",
	}

	d := []Blog{b1, b2, b3}

	w.Header().Set("Content-Type", "text/html")
	err := homeView.Template.ExecuteTemplate(w, homeView.Layout, d)
	if err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {


	w.Header().Set("Content-Type", "text/html")
	err := contactView.Template.ExecuteTemplate(w, contactView.Layout, nil)
	if err != nil {
		panic(err)
	}
}

func main() {

	db, err := gorm.Open("mysql", "root:password@/adsTest")
	if err!= nil {
		panic(err.Error())
	}
	defer db.Close()
	db.LogMode(true)

	db.AutoMigrate(&Blog{})

	homeView = views.NewView("main","views/home.gohtml")
	contactView = views.NewView("main", "views/contact.gohtml")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":3040", r)
}