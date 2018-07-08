package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:dds3@1@/cursogo")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	stmt, _ := db.Prepare("delete from usuarios where id = ?")

	stmt.Exec(5)

}
