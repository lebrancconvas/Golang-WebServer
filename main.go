package main

import (
	"fmt"
	"strings"
	"net/http"
)

func main() {
	userDB := map[string]int{
		"tarwaan":21,
		"mint":20,
		"keyboard":17,
	}

	http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request) {
		fmt.Fprintf(w,"Hello Golang Server")
	})

	// Route
	http.HandleFunc("/user/",func (w http.ResponseWriter,r *http.Request) {
		name := r.URL.Path[len("/user/"):]
		age := userDB[name]
		fmt.Fprintf(w,"I am %s , My Age is %d",strings.Title(name),age)
	})
	
	http.HandleFunc("/information",func(w http.ResponseWriter,r *http.Request) {
		fmt.Fprintf(w,"More Information Contact : Scarletto Studio")
	})
	
	http.HandleFunc("/hello",func(w http.ResponseWriter,r *http.Request) {
		fmt.Fprintf(w,"Hello World")
	})
	
	http.HandleFunc("/wtf",func(w http.ResponseWriter,r *http.Request) {
		fmt.Fprintf(w,"What The Fuck.")
	})
	
	http.ListenAndServe(":8080",nil)
}