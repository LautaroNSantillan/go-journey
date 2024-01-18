package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/LautaroNSantillan/my-go-journey/tree/mysql-books/pkg/models"
	"github.com/LautaroNSantillan/my-go-journey/tree/mysql-books/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBooks(res http.ResponseWriter, req *http.Request) {
	newBooks := models.GetAllBooks()
	jsonRes, _ := json.Marshal(newBooks)
	res.Header().Set("Content-Type", "pkglication/json")
	res.WriteHeader(http.StatusOK)
	res.Write(jsonRes)
}

func GetBookById(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("parsing error")
	}

	bookByID, _ := models.GetBookById(ID)

	jsonRes, _ := json.Marshal(bookByID)
	res.Header().Set("Content-Type", "pkglication/json")
	res.WriteHeader(http.StatusOK)
	res.Write(jsonRes)
}

func CreateBook(res http.ResponseWriter, req *http.Request) {
	CreatedBook := models.Book{}

	utils.ParseBody(req, CreatedBook)
	dbBook := CreatedBook.CreateBook()

	jsonRes, _ := json.Marshal(dbBook)
	res.WriteHeader(http.StatusCreated)
	res.Write(jsonRes)
}

func DeleteBook(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("parsing error")
	}

	delBook := models.DeleteBook(ID)

	jsonRes, _ := json.Marshal(delBook)
	res.Header().Set("Content-Type", "pkglication/json")
	res.WriteHeader(http.StatusOK)
	res.Write(jsonRes)
}

func UpdateBook(res http.ResponseWriter, req *http.Request) {
	var updatedBook = &models.Book{}

	utils.ParseBody(req, updatedBook)

	vars := mux.Vars(req)
	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("parsing error")
	}

	dbBook, db := models.GetBookById(ID)

	if updatedBook.Name != "" {
		dbBook.Name = updatedBook.Name
	}
	if updatedBook.Author != "" {
		dbBook.Author = updatedBook.Author
	}
	if updatedBook.Publication != "" {
		dbBook.Publication = updatedBook.Publication
	}

	db.Save(&dbBook)

	jsonRes, _ := json.Marshal(dbBook)
	res.Header().Set("Content-Type", "pkglication/json")
	res.WriteHeader(http.StatusOK)
	res.Write(jsonRes)
}
