package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"jakataGo/controllers"
	"log"
	"net/http"
	"os"
)

type Review struct {
	gorm.Model
	Review string
	Client string
	Stylist string
}

type BlogPost struct {
	gorm.Model
	Title string
	Slug string
	Author string
	MetaImg string
	Publish int
	Para string
}

type BlogPara struct {
	gorm.Model
	BlogId int
	Para string
	ParaPic string
	ParaPicAlt string
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

func dbConn() (db *gorm.DB) {
	dbhost     := os.Getenv("DB_HOST")
	dbport     := os.Getenv("DB_PORT")
	dbuser     := os.Getenv("DB_USER")
	dbpassword := os.Getenv("DB_PASSWORD")
	dbname     := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbhost, dbport, dbuser, dbpassword, dbname)

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	db := dbConn()
	db.AutoMigrate(&Review{}, &BlogPost{}, &BlogPara{})
	db.Close()
	db.LogMode(true)

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


	http.ListenAndServe(":" + port, r)

}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/public/images/favicon.ico")
}

func reviews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := dbConn()

	revs := []Review{}
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

	blogs := []BlogPost{}
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
