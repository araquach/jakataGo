package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

var ErrNoRecord  = errors.New("Models: No matching record found")

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
