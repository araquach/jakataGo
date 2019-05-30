package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

const (
	user = "root"
	password = "password"
	dbname = "adsTest"
)

type User struct {
	gorm.Model
	Name string
	Address string
}

func main(){
	psqlInfo := fmt.Sprintf("%s:%s@/%s?parseTime=true", user, password, dbname)
	db, err := gorm.Open("mysql", psqlInfo)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Connection Established")
	defer db.Close()

	db.LogMode(true)

	db.Debug().DropTableIfExists(&User{})
	db.Debug().AutoMigrate(&User{})

	user := &User{Name: "Adam", Address: "32 Glossop Close"}
	db.Create(user)

	var users []User = []User{
		User{Name: "John", Address: "At Home"},
		User{Name: "Pete", Address: "At Work"},
		User{Name: "Roger", Address: "54 Awful Space"},
	}

	for _, user := range users {
		db.Create(&user)
	}

	db.Table("users").Where("address = ?", "At Work").Update("name", "Liz")

	db.Table("users").Where("id= ?", 1).Delete(&User{})

	db.Where("name=?", "John").Delete(&User{})

	db.Debug().Where("name LIKE ?", "%ti%").Find(&user)
}