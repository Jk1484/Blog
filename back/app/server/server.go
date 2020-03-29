package server

import (
	"blog/back/app/router"
	"blog/back/config"
	"net/http"
)

// Init initializes server
func Init() (err error) {
	var (
		cfg = config.Peek().Server
		srv *http.Server
	)

	srv = &http.Server{
		Addr:    cfg.Port,
		Handler: router.Get(),
	}

	return srv.ListenAndServe()
}
