package app

import (
	"io"
	"net/http"
	"text/template"
)

func (s *server) routes() {
	s.Router.HandleFunc("/api/v1/", s.handleAPI())
	s.Router.HandleFunc("/transfer", s.handleTransfer())
	s.Router.HandleFunc("/", s.handleIndex())
}

func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
	}
}

func (s *server) handleTransfer() http.HandlerFunc {
	// transfer should get some url query parameters
	// this client is here emulating a database query
	c := newClient("Jon Doe", "jondoe@mail.com", "Jon Doe's street")
	t, err := template.ParseFiles("./transfer.gohtml")
	if err != nil {
		panic(err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		err := t.Execute(w, c)
		if err != nil {
			panic(err)
		}
	}
}

// TODO: handle the API route!
func (s *server) handleAPI() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func newClient(name, email, address string) *Client {
	return &Client{
		name,
		email,
		address,
	}
}
