package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func main() {
	db,err := sql.Open("mysql","root:@/idols")

	if err != nil{
		fmt.Println(err)
	}
	defer db.Close()

	rows,err := db.Query("SELECT id,Name,Age,Generation FROM bnk48")
	if err != nil{
		fmt.Println(err)
	}

	for rows.Next(){
		var id int
		var name string
		var age int
		var gen int
		err = rows.Scan(&id,&name,&age,&gen)
		fmt.Printf("ID : %d\nName : %s\nAge : %d\nGeneration : %d\n\n",id,name,age,gen)
	}
}