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

	// fetch desde la base de datos
	/*
		var (
			id   int
			name string
		)
		rows, err := db.Query("select id, fullname from users where id = ?", 1)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&id, &name)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Hola, :D", id, name)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	*/

	// preparar sql para usar multiples veces con parametros placeholders, en mysql es ?
	/*
		var (
			id       int
			fullname string
		)
		stmt, err := db.Prepare("select id, fullname from users where id = ?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		rows, err := stmt.Query(1)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&id, &fullname)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Hola, :D", id, fullname)
		}
		if err = rows.Err(); err != nil {
			log.Fatal(err)
		}
	*/
	/*
		stmt, err := db.Prepare("select fullname from users where id = ?")
		if err != nil {
			log.Fatal(err)
		}
		var name string
		err = stmt.QueryRow(1).Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(name)*/
}
