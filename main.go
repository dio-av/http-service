package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type client struct {
	Name    string
	Email   string
	Address string
}

func newClient(name, email, address string) client {
	return client{
		name,
		email,
		address,
	}
}

func main() {
	const tmpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>Transactions</title>
	</head>
	<body>
		Hello {{ .Name }}. Your email is {{ .Email }} and your password is: {{ .Address }}		
	</body>
</html>`

	c := newClient("Jon Doe", "jondoe@mail.com", "Jon Doe's street")

	t, err := template.ParseFiles("./transfer.gohtml")
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		io.WriteString(w, `<doctype html>
			<html>
				<head>
					<title>Gopay</title>
				</head>
				<body>
					<p>Hello! Welcome to Gopay.</p>
				</body>
			</html>`,
		)
	})
	router.HandleFunc("/transfer", func(w http.ResponseWriter, r *http.Request) {
		err = t.Execute(w, c)
		if err != nil {
			panic(err)
		}
	})
	srv := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
