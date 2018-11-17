package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/",index)
	http.HandleFunc("/login",login)
	http.HandleFunc("/signup",signup)
	http.HandleFunc("/single",single)
	http.HandleFunc("/loginpage",loginpage)
	http.ListenAndServe(":8080",nil)
}

func index(w http.ResponseWriter,r *http.Request) {
	http.ServeFile(w,r,"index.html")
}

func login(w http.ResponseWriter,r *http.Request) {
	http.ServeFile(w,r,"login.html")
}

func signup(w http.ResponseWriter,r *http.Request) {
	http.ServeFile(w,r,"signup.html")
}

func single(w http.ResponseWriter,r *http.Request) {
	http.ServeFile(w,r,"single.html")
}

func loginpage(w http.ResponseWriter,r *http.Request) {
	fmt.Println("Method : ",r.Method)
	r.ParseForm()
	fmt.Println("Username : ",r.Form["username"])
	fmt.Println("Password : ",r.Form["password"])
	http.ServeFile(w,r,"loginpage.html")
} 