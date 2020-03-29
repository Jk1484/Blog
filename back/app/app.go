package app

import (
	"blog/back/app/db"
	"blog/back/app/router"
	"blog/back/app/server"
	"log"
	"os"
)

const file = "./app.log"

func init() {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(f)
	}
	log.SetOutput(f)
}

// Run runs site
func Run() {
	var (
		err error
	)

	err = db.Init()
	if err != nil {
		log.Panicln("Database err:", err)
		return
	}
	defer db.Exit()

	err = router.Init()
	if err != nil {
		log.Panicln("Router err:", err)
		return
	}

	err = server.Init()
	if err != nil {
		log.Println("Server err:", err)
		return
	}
}
