package models

import (
	"github.com/LautaroNSantillan/my-go-journey/tree/mysql-books/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `gorm:"author"`
	Publication string `gorm:"publish_date"`
}

func initDB() {
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
	var Books []Book
	db.Find(&Books)
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
