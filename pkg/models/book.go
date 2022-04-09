package models

import (
	"github.com/RakibSiddiquee/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

// Book is a book struct
type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// init connects db, get db and runs auto migrations
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// CreateBook insert a new book
func (b *Book) CreateBook() *Book {
	//db.NewRecord(b)
	db.Create(&b)
	return b
}

// GetAllBooks fetch all the books
func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

// GetBookById gets book by ID
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

// DeleteBook deletes a book by ID
func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(&book)
	return book
}
