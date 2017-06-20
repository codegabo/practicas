package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// "mysql" es un alias para el paquete github.com/go-sql-driver/mysql
	db, err := sql.Open("mysql",
		"ehnqhlcm_juangabrielmm:31GaD3y#il06er%^7s8@tcp(mysql.anaxanet.com:3306)/ehnqhlcm_edcomments")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Statements that Modify Data
	// Use Exec(), preferably with a prepared statement,
	// to accomplish an INSERT, UPDATE, DELETE, or other statement that doesn’t return rows.
	// The following example shows how to insert a row and inspect metadata about the operation:
	stmt, err := db.Prepare("INSERT INTO users(username, email, fullname, password) VALUES(?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec("withoutsioul", "juioan@g", "juan gabrielll", "123")
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("usuario creado con exito\nID = %d, affected = %d\n", lastId, rowCnt)

}
