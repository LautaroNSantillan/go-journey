package main

import (
	"fmt"
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
	http.HandleFunc("/add-film/", addFilm)

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

func addFilm(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	director := r.PostFormValue("director")

	htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'>%s - %s</li>", title, director)

	tmpl, _ := template.New("myTempl").Parse(htmlStr)

	tmpl.Execute(w, nil)

}
