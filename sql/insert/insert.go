package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:dds3@1@/cursogo")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	tx, _ := db.Begin()

	stmt, _ := db.Prepare("insert into usuarios(id,nome) values(?,?)")

	_, err = stmt.Exec(3, "Mariangela")
	stmt.Exec(4, "César")
	stmt.Exec(5, "Laizé")

	if err != nil {
		tx.Rollback()
		log.Fatal(err)

	}

	tx.Commit()

}
