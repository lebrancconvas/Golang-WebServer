package main

import (
	"fmt"
	"strings"
	"net/http"
	"github.com/gorilla/mux"
	"html/template"
)

type Single struct{
	Name string
	Album string
}

func main() {
	var templates = template.Must(template.ParseFiles("single.html"))

	router := mux.NewRouter()

	idolDB := map[string]int{
		"may":12,
		"daoruang":23,
		"joanne":21,
		"memechan":25,
		"tarwaan":21,
		"mint":20,
		"keyboard":17,
	}

	router.HandleFunc("/",func(w http.ResponseWriter,r *http.Request) {
		http.ServeFile(w,r,"index.html")
	})

	router.HandleFunc("/single",func(w http.ResponseWriter,r *http.Request) {
		mySingle := Single{"Milky Way","Party Time"}
		templates.ExecuteTemplate(w,"single.html",mySingle)
	})

	router.HandleFunc("/information",func(w http.ResponseWriter,r *http.Request) {
		fmt.Fprintf(w,"More Information Contact : Scarletto Studio")
	})
	
	router.HandleFunc("/hello",func(w http.ResponseWriter,r *http.Request) {
		fmt.Fprintf(w,"Hello World")
	})
	
	router.HandleFunc("/wtf",func(w http.ResponseWriter,r *http.Request) {
		fmt.Fprintf(w,"What The Fuck.")
	})

	router.HandleFunc("/idol",func(w http.ResponseWriter,r *http.Request) {
		http.ServeFile(w,r,"idol.html")
	})

	router.HandleFunc("/idol/{name}",func(w http.ResponseWriter,r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		age := idolDB[name]
		fmt.Fprintf(w,"My name is %s , My age is %d",strings.Title(name),age)
	}).Methods("GET")

	http.ListenAndServe(":8080",router)
}