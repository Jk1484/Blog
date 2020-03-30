package db

import (
	"blog/back/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var client *sql.DB

// Init initialises connection with database
func Init() (err error) {
	var (
		cfg = config.Peek().Database
	)

	psqlInfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable connect_timeout=10",
		cfg.User, cfg.Password, cfg.DBName, cfg.Host, cfg.Port)

	client, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		return
	}

	err = client.Ping()

	if err != nil {
		return
	}

	log.Printf("database: established with %s:%s/%s\n", cfg.Host, cfg.Port, cfg.DBName)

	return
}

// Exit closes connection to database
func Exit() {

}

func Client() *sql.DB {
	return client
}
