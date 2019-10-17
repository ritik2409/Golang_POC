package migrations

import (
	"log"
	obj "project/dto"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Gorm_migrations() {

	db, err := gorm.Open("mysql", "root:root123@tcp(127.0.0.1:3306)/golang_poc")
	defer db.Close()
	if err != nil {
		log.Println("Connection Failed to Open")
	}
	log.Println("Connection Established")

	//  db.CreateTable(&obj.Users{})
	// db.Debug().DropTableIfExists(&obj.Users{})
	db.Debug().AutoMigrate(&obj.Users{})
	// log.Println("Ho raha hai table creation!!")

}
