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
	var (
		id   int
		name string
	)
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
			user_id int
			id      int
			content string
		)
		stmt, err := db.Prepare("select user_id, id, content from comments where id = ?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		rows, err := stmt.Query(12)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			err = rows.Scan(&user_id, &id, &content)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Hola, tu usuario, id y contenido son:\n ", user_id, id, content)
		}
		if err = rows.Err(); err != nil {
			log.Fatal(err)
		}
	*/

	// preparar la declaracion sql para ser usada con parametros placeholder
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
