package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	// XXX make doc directory be an arg
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)
	http.HandleFunc("/page", fill)

	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

type Page struct {
	Title string
	Body  string
}

func fill(w http.ResponseWriter, r *http.Request) {
	page := Page{
		Title: "My Page",
		Body:  "Hello, World!",
	}
	tmpl := template.Must(template.ParseFiles("page.tmpl"))
	tmpl.Execute(w, page)
}
