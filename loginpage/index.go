package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/",login)
	http.HandleFunc("/login",loginpage)
	http.ListenAndServe(":8080",nil)
}

func login(w http.ResponseWriter,r *http.Request) {
	http.ServeFile(w,r,"login.html")
}

func loginpage(w http.ResponseWriter,r *http.Request) {
	fmt.Println("Method : ",r.Method)
} 