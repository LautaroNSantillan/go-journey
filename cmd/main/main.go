package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LautaroNSantillan/my-go-journey/tree/mysql-books/pkg/models"
	"github.com/LautaroNSantillan/my-go-journey/tree/mysql-books/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	models.InitDB()
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)

	http.Handle("/", router)
	fmt.Printf("Starting server on port 8081\n")
	log.Fatal(http.ListenAndServe("localhost:8081", router))
}
