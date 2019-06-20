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
	Para []string
	Link string
	Author string
}

type TeamMember struct {
	Fname string
	Lname string
	Image string
	Level string
	Salon int
	Para1 string
	Para2 string
	Para3 string
	FavStyle string
	FavProduct string
	Price int
	ReviewLink string
	Class string
	Position int
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
		r.HandleFunc("/api/team", team).Methods("GET")


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


	http.HandleFunc("/favicon.ico", faviconHandler)


	http.ListenAndServe(":5090", r)

}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/public/images/favicon.ico")
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
		"http://via.placeholder.com/1000x1000",
		"blog 1 pic",
		"Blog Post One",
		[]string{"First para of Blog 1", "Second Para of Blog 1", "Third Para of Blog One"},
		"/blog/1",
		"Adam",
	}
	b2 := BlogPost{
		"blog-two",
		"http://via.placeholder.com/1000x1000",
		"blog 2 pic",
		"Blog Post Two",
		[]string{"First para of Blog 2", "Second Para of Blog 2", "Third Para of Blog Two"},
		"/blog/2",
		"Nat",
	}
	b3 := BlogPost{
		"blog-three",
		"http://via.placeholder.com/1000x1000",
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

func team(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	t1 := TeamMember{
		"Adam",
		"Carter",
		"http://via.placeholder.com/1000x1000",
		"Director",
		1,
		"Adam is a great stylist",
		"He's the owner",
		"Make sure you book in!",
		"Short Bobs",
		"Anti Gravity",
		150,
		"/reviews/adam",
		"adam",
		1,
	}
	t2 := TeamMember{
		"Jimmy",
		"Sharpe",
		"http://via.placeholder.com/1000x1000",
		"Director",
		1,
		"Jim is a great stylist",
		"He's the manager",
		"He's a great stylist!",
		"Bold short cuts",
		"Mess Up",
		140,
		"/reviews/jimmy",
		"jimmy",
		2,
	}
	t3 := TeamMember{
		"Natalie",
		"Doxey",
		"http://via.placeholder.com/1000x1000",
		"Freelance Senior Stylist",
		1,
		"Nat is a great stylist",
		"She's freelance",
		"She's great at extensions",
		"Crazy Colours",
		"Blow Me",
		140,
		"/reviews/nat",
		"nat",
		3,
	}

	team := []TeamMember{t1, t2, t3}

	json, err := json.Marshal(team)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}
