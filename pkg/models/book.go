package models

import (
	"fmt"

	"github.com/LautaroNSantillan/my-go-journey/tree/mysql-books/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func InitDB() {
	config.ConnectToDB()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (book *Book) CreateBook() *Book {
	db.NewRecord(book)
	db.Create(&book)
	return book
}

func GetAllBooks() []Book {
	if db == nil {
		fmt.Println("conected")
		config.ConnectToDB() // Ensure the database is connected
	}

	var Books []Book
	if db != nil {
		db.Find(&Books)
		fmt.Println("found book")
	}
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var theBook Book
	db := db.Where("ID=?", Id).Find(&theBook)
	return &theBook, db
}

func DeleteBook(ID int64) Book {
	var delBook Book
	db.Where("ID=?", ID).Delete(delBook)
	return delBook
}
