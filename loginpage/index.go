package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
	"time"
)

type Cookie struct{
	Name string
	Value string
	Expires time.Time
}

func main() {
	http.HandleFunc("/",index)
	http.HandleFunc("/login",login)
	http.HandleFunc("/signup",signup)
	http.HandleFunc("/single",single)
	http.HandleFunc("/loginpage",loginpage)
	http.HandleFunc("/cookie",cookie)
	http.HandleFunc("/upload",upload)
	http.HandleFunc("/uploadpage",uploadpage)
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

func cookie(w http.ResponseWriter,r *http.Request) {
	expiration := time.Now().Add(time.Hour*24*365)
	cookie := http.Cookie{Name : "user",Value : "HelloGolang",Expires : expiration}
	http.SetCookie(w,&cookie)
	fmt.Fprintf(w,"Create Cookie.")
}

func upload(w http.ResponseWriter,r *http.Request) {
	http.ServeFile(w,r,"upload.html")
}

func uploadpage(w http.ResponseWriter,r *http.Request) {
	// รับ File ที่ต้องการ Upload
	file,handle,err := r.FormFile("file")
	defer file.Close()
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w,"%v",handle.Header)

	// บอกตำแหน่งว่าต้องการ Upload File ที่เราเลือกไปไว้ที่ตรงไหน
	f,err := os.OpenFile("./img/"+handle.Filename,os.O_CREATE,0666)
	//แสดงข้อความเมื่อการ Upload ล้มเหลว
	if err != nil{
		fmt.Println(err)
		return
	}
	defer f.Close()
	
	// ย้าย File ที่อัพโหลดไปไว้บนตำแหน่งที่เลือก
	io.Copy(f,file)
	// แสดงข้อความเมื่อการ Upload สำเร็จ
	fmt.Fprintf(w,"\nUpload Success!!")
}

func loginpage(w http.ResponseWriter,r *http.Request) {
	fmt.Println("Method : ",r.Method)
	r.ParseForm()
	fmt.Println("Username : ",r.Form["username"])
	fmt.Println("Password : ",r.Form["password"])
	http.ServeFile(w,r,"loginpage.html")
} 