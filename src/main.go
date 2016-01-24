package main

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	db, err := gorm.Open("postgres", "user=gorm password=gorm dbname=gorm sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	db.DropTableIfExists(&Owner{}, &Book{}, &Author{})
	db.CreateTable(&Owner{}, &Book{}, &Author{})

}

type Owner struct {
	gorm.Model
	FirstName string
	LastName  string
	Books     []Book
}

type Book struct {
	gorm.Model
	Name        string
	PublishDate time.Time
	OwnerID     uint     `sql:"index"`
	Authors     []Author `gorm:"many2many:books_authors"`
}

type Author struct {
	gorm.Model
	FirstName string
	LastName  string
}
