package main

import (
	"log"
	"net/http"
	"text/template"
)

type Film struct {
	Title    string
	Director string
}

func main() {

	http.HandleFunc("/", handler1)

	log.Fatal(http.ListenAndServe(":8000", nil))

}

func handler1(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFiles("index.html"))

	films := map[string][]Film{
		"Films": {
			{Title: "1", Director: "1"},
			{Title: "2", Director: "2"},
			{Title: "3", Director: "3"},
		},
	}
	templ.Execute(w, films)
}
