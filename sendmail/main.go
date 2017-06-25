package main

import (
	"fmt"
	"log"
	// bytes permite escribir información, en vez de concatenar strings
	"bytes"
	"net/mail"
	// net/smtp permite enviar el correo
	"net/smtp"
	// crypto permite configurar las opciones de tls de gmail
	"crypto/tls"
	"html/template"
)

type Dest struct {
	Name string
}

func checkErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
func main() {
	from := mail.Address{"Appqui Colombia", "ggabooo97@gmail.com"}
	to := mail.Address{"Gabriel Mogollón", "ggabooo@hotmail.com"}
	subject := "Enviando correo desde golang"
	//destinatario
	dest := Dest{Name: to.Address}

	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject
	headers["Content-Type"] = `text/html; charset="UTF -8"`

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	// pasarle la informacion al template
	t, err := template.ParseFiles("template.html")
	checkErr(err)
	buf := new(bytes.Buffer)
	err = t.Execute(buf, dest)
	checkErr(err)
	// enviar el contenido del template con el destinatario
	message += buf.String()

	servername := "smtp.gmail.com:465"
	host := "smtp.gmail.com"
	// autentificacion
	auth := smtp.PlainAuth("", "ggabooo97@gmail.com", "3108100788quatrimoto1", host)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	conn, err := tls.Dial("tcp", servername, tlsConfig)
	checkErr(err)
	// cliente de la cone4xion
	client, err := smtp.NewClient(conn, host)
	checkErr(err)

	err = client.Auth(auth)
	checkErr(err)

	err = client.Mail(from.Address)
	checkErr(err)

	err = client.Rcpt(to.Address)
	checkErr(err)

	w, err := client.Data()
	checkErr(err)

	_, err = w.Write([]byte(message))
	checkErr(err)

	err = w.Close()
	checkErr(err)

	client.Quit()
}
