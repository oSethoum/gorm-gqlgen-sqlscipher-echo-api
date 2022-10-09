package db

import (
	"log"

	sqlite "github.com/oSethoum/sqlite3"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {

	db, err := gorm.Open(sqlite.Open("db.sqlite?_fk=1&_pragma_key=password"), &gorm.Config{
		Logger: nil,
	})

	if err != nil {
		log.Fatal(err)
	}

	DB = db
}
