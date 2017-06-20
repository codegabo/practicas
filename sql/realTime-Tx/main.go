package main

import (
	"database/sql"
	"log"
"fmt"
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


  var content string
  err = db.QueryRow("select content from comments").Scan(&content)
  if err != nil {
  	log.Fatal(err)
  }
  fmt.Println(content)


}
