package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"jakataGo/controllers"
	"log"
	"net/http"
)

type Review struct {
	Review string
	Client string
	Stylist string
}

type BlogPost struct {
	Slug string
	Img string
	Alt string
	Title string
	Paras []string
	Link string
	Author string
}

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
		// r.Handle("/prices", pageC.PricesView).Methods("GET")
		r.Handle("/recruitment", pageC.RecruitmentView).Methods("GET")
		r.HandleFunc("/api/reviews", reviews).Methods("GET")
		r.HandleFunc("/api/blogs", blogs).Methods("GET")


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

	http.ListenAndServe(":5090", r)

}

func reviews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	r1 := Review{
		"Great job!",
		"Anna Alexander",
		"Nat",
	}
	r2 := Review{
		"Awsome!",
		"Jackie Alexander",
		"Adam",
	}
	revs := []Review{r1, r2}

	json, err := json.Marshal(revs)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}

func blogs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	b1 := BlogPost{
		"blog-one",
		"/blog/1",
		"blog 1 pic",
		"Blog Post One",
		[]string{"First para of Blog 1", "Second Para of Blog 1", "Third Para of Blog One"},
		"/blog/1",
		"Adam",
	}
	b2 := BlogPost{
		"blog-two",
		"/blog/2",
		"blog 2 pic",
		"Blog Post Two",
		[]string{"First para of Blog 2", "Second Para of Blog 2", "Third Para of Blog Two"},
		"/blog/2",
		"Nat",
	}
	b3 := BlogPost{
		"blog-three",
		"/blog/3",
		"blog 3 pic",
		"Blog Post Three",
		[]string{"First para of Blog 3", "Second Para of Blog 3", "Third Para of Blog Three"},
		"/blog/1",
		"Adam",
	}

	blogs := []BlogPost{b1, b2, b3}

	json, err := json.Marshal(blogs)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}
