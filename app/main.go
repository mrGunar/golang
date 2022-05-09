package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	name       string
	age        uint16
	balance    int16
	avg_grades float64
}

func home_page(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("app/index.html")
	user := User{"Alex", 30, 5000, 10.0}
	tmpl.Execute(w, user)
}

func about_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About myself")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/about", about_page)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
