package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LautaroNSantillan/my-go-journey/tree/mysql-books/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)

	http.Handle("/", router)
	fmt.Printf("Starting server on port 8080\n")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
