package app

import (
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type server struct {
	db     *gorm.DB
	Router *chi.Mux
}

func NewServer(database *gorm.DB) *server {
	s := &server{
		db: database,
	}
	s.routes()

	return s
}

type Client struct {
	Name    string
	Email   string
	Address string
}
