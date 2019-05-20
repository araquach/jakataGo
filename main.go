package main

import (
	"github.com/gorilla/mux"
	"jakataGo/controllers"
	"net/http"
)

func main() {
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

	http.ListenAndServe(":8080", r)

}
