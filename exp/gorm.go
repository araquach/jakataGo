package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	user = "root"
	password = "password"
	dbname = "adsTest"
)

type User struct {
	gorm.Model
	Name string
	Email string `gorm:"not null;unique_index"`
	Orders []Order
}

type Order struct {
	gorm.Model
	UserID uint
	Amount int
	Description string
}

func main() {
	psqlInfo := fmt.Sprintf("%s:%s@/%s?parseTime=true", user, password, dbname)
	db, err := gorm.Open("mysql", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.LogMode(true)

	db.AutoMigrate(&User{}, &Order{})

	var user User
	user.Email = "ara@yah.com"
	db.Preload("Orders").Where(user).First(&user)
	if db.Error != nil {
		panic(db.Error)
	}

	fmt.Println("Email:", user.Email)
	fmt.Println("Number of orders:", len(user.Orders))
	fmt.Println("Orders:", user.Orders)



	//createOrder(db, user, 1001, "Fake Description #1")
	//createOrder(db, user, 9999, "Fake Description #2")
	//createOrder(db, user, 8800, "Fake Description #3")


	// Using WHERE method
	/*
		var u User
		maxId := 3

		db.Where("id > ?", maxId).First(&u)
		if db.Error != nil {
			panic(db.Error)
		}
		fmt.Println(u)
	*/

	// Getting first user
	/*
		var u User
		db.First(&u)
		if db.Error != nil {
			panic(db.Error)
		}
		fmt.Println(u)
	*/

	// Creating a user
	/*
		name, email := getInfo()
		u := &User{
			Name: name,
			Email: email,
		}

		if err = db.Create(u).Error
		err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", u)
	*/

	// Querying an existing user by email
	/*
			var u User
			u.Email = "araquach@yahoo.co.uk"

			db.Where(u).First(&u)
			if db.Error != nil {
				panic(db.Error)
			}
			fmt.Println(u)
		}
	*/

	// Querying for multiple users
	/*
		var users []User
		db.Find(&users)
		if db.Error != nil {
			panic(db.Error)
		}
		fmt.Println("Retrieved", len(users), "users.")
		fmt.Println(users)

	*/

	//
}

//func createOrder(db *gorm.DB, user User, amount int, desc string) {
//	db.Create(&Order{
//		UserID: user.ID,
//		Amount: amount,
//		Description: desc,
//	})
//	if db.Error != nil {
//		panic(db.Error)
//	}
//}


// Function to generate new user
/*
	func getInfo() (name, email string) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("What is your name?")
		name, _ = reader.ReadString('\n')
		name = strings.TrimSpace(name)
		fmt.Println("What is your email?")
		email, _ = reader.ReadString('\n')
		email = strings.TrimSpace(email)
		return name, email
	}

*/

