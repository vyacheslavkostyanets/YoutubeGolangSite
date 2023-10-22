package main

import (
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8080", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	view, _ := template.ParseFiles("index.html")
	data := "HELLO WORLD"
	view.Execute(w, data)
}
