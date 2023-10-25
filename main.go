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
	house := map[string]int{
		"floors":  10,
		"windows": 1000,
		"rooms":   100,
	}
	personOne := Person{
		Name:   "Slava",
		Age:    30,
		Health: true,
	}

	alldata := map[string]interface{}{
		"House":  house,
		"Person": personOne,
	}

	view.Execute(w, alldata)
}

type Person struct {
	Name   string
	Age    int
	Health bool
}
