package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request) {
		fmt.Fprintf(w,"Hello Golang Server")
	})
	http.HandleFunc("/index",func(w http.ResponseWriter,r *http.Request) {
		fmt.Fprintf(w,"Index Body")
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
	fmt.Println("Go Server Run at :8080")
}