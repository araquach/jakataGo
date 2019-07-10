package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
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
	gorm.Model
	Name string
	Salon int
	Level string
	Image string
	Para1 string
	Para2 string
	Para3 string
	FavStyle string
	FavProd string
	Price string
	ReviewLink string
	Class string
	Position string
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
	db.AutoMigrate(&Review{}, &BlogPost{}, &BlogPara{}, &TeamMember{})
	db.Close()
	db.
	db.LogMode(true)

	srv := &http.Server{
		Addr: ":" + port,
		Handler: routes(),
	}

	log.Printf("Starting server on %s", port)
	err := srv.ListenAndServe()
	log.Fatal(err)
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

	db := dbConn()
	team := []TeamMember{}
	db.Where("Salon = 1").Order("position asc").Find(&team)
	db.Close()

	json, err := json.Marshal(team)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}
