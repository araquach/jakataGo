package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"jakataGo/controllers"
	"net/http"

)

func main() {
	db, err := sql.Open("mysql", "root:password@/jakataGo")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	pageC := controllers.NewPage()

	r := mux.NewRouter()
	r.Handle("/", pageC.HomeView).Methods("GET")
	r.Handle("/team", pageC.TeamView).Methods("GET")
	r.Handle("/team_ind", pageC.TeamIndView).Methods("GET")
	r.Handle("/blog", pageC.BlogView).Methods("GET")
	r.Handle("/blog_ind", pageC.BlogIndView).Methods("GET")
	r.Handle("/details", pageC.DetailsView).Methods("GET")
	r.Handle("/contact", pageC.ContactView).Methods("GET")
	r.Handle("/reviews", pageC.ReviewsView).Methods("GET")
	r.Handle("/salon", pageC.SalonView).Methods("GET")
	r.Handle("/prices", pageC.PricesView).Methods("GET")
	r.Handle("/recruitment", pageC.RecruitmentView).Methods("GET")

	// Styles
	assetHandler := http.FileServer(http.Dir("./public/css/"))
	assetHandler = http.StripPrefix("/public/css/", assetHandler)
	r.PathPrefix("/public/css/").Handler(assetHandler)

	// JS
	jsHandler := http.FileServer(http.Dir("./public/js/"))
	jsHandler = http.StripPrefix("/public/js/", jsHandler)
	r.PathPrefix("/public/js/").Handler(jsHandler)

	//Images
	imageHandler := http.FileServer(http.Dir("./public/images/"))
	r.PathPrefix("/images/").Handler(http.StripPrefix("/images/", imageHandler))

	http.ListenAndServe(":8080", r)

}
