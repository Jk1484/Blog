package router

import (
	"blog/back/app/handlers"
	"blog/back/config"
	"log"

	"github.com/julienschmidt/httprouter"
)

var router *httprouter.Router

// Init initializes router
func Init() (err error) {
	var (
		cfg = config.Peek().Server
	)

	router = httprouter.New()

	router.GET("/", handlers.Health)

	log.Println("Router initialized with port ", cfg.Port)

	return
}

// Get returns router
func Get() *httprouter.Router {
	return router
}
