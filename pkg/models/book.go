package models

import (
	"mySQL-Bookstore-Management/pkg/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Author      string `json:"author"`
	Name        string `json:"name"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetBookById(id int32) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", id).Find(&book)
	return &book, db
}

func DeleteBook(id int32) Book {
	var book Book
	db.Where("ID=?", id).Delete(book)
	return book
}
