package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"os"
)

func main(){
	db, err := sql.Open("mysql", "root:19850922@/world")
	if(err != nil){
		os.Exit(1)
	}
	// fmt.Printf("%t\n", db)
	
	row := db.QueryRow("select ID, NAME from city limit 1")
	var id int
	var name string
	row.Scan(&id, &name)
	fmt.Printf("%d %s\n", id, name)
}