package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

type Book struct {
	Title  string
	Author string
}

func main() {
	fileServer := http.FileServer((http.Dir("template1/index.html")))
	fmt.Println(fileServer)
	http.Handle("/", http.HandlerFunc(ShowBooks))
	http.Handle("/delete", http.HandlerFunc(DeleteBook))
	http.ListenAndServe(":8080", nil)
}

func ShowBooks(w http.ResponseWriter, r *http.Request) {
	book := Book{"Building Web Apps with Go", "Jeremy Saenz"}

	fp := path.Join("template1", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
