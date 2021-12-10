package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Entities struct {
	id               uint16
	name             string
	amount           uint16
	inventory_number uint16
	department_name  string
}

func main_page(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/main_page.html")
	tmpl.Execute(w, "templates/main_page.html")
}

func login(page http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(page, "Log in")
}

func choose(page http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(page, "Input Output Delete Update")
}

func input(page http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(page, "Input")
}

func output(page http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(page, "Output")
}

func delete(page http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(page, "Delete")
}

func update(page http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(page, "Update")
}

func logout(page http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(page, "Log out")
}

func HandlRequest() {
	http.HandleFunc("/", main_page)
	http.HandleFunc("/login", login)
	http.HandleFunc("/choose", choose)
	http.HandleFunc("/input", input)
	http.HandleFunc("/output", output)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/update", update)
	http.HandleFunc("/logout", logout)
	http.ListenAndServe(":8080", nil)
}

func main() {
	HandlRequest()
}
