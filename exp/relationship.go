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

type Place struct {
	gorm.Model
	Name string
	Town Town
	TownId int `gorm:"ForeignKey:id"`
}

type Town struct {
	gorm.Model
	Name string
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

	db.DropTableIfExists(&Place{}, &Town{})
	db.AutoMigrate(&Place{}, &Town{})
	db.Model(&Place{}).AddForeignKey("town_id", "towns(id)", "CASCADE", "CASCADE")

	t1 := Town{
		Name: "Warrington",
	}
	t2 := Town{
		Name: "Manchester",
	}
	t3 := Town{
		Name: "Liverpool",
	}

	p1 := Place{
		Name: "Stockton Heath",
		Town: t1,
	}
	p2 := Place{
		Name: "Didsbury",
		Town: t2,
	}
	p3 := Place{
		Name: "Everton",
		Town: t3,
	}

	db.Save(&p1)
	db.Save(&p2)
	db.Save(&p3)

	fmt.Println("t1==>", t1.Name, "p1==>", p1.Name)
	fmt.Println("t2==>", t2.Name, "p2s==>", p2.Name)
	fmt.Println("t2==>", t3.Name, "p2s==>", p3.Name)

	db.Where("name=?", "Warrington").Delete(&Town{})

}
