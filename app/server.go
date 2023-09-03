package app

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type server struct {
	db     *gorm.DB
	router *mux.Router
	//email  EmailSender
}

type Client struct {
	Name    string
	Email   string
	Address string
}
