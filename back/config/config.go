package config

import (
	"encoding/json"
	"log"
	"os"
)

var (
	configPath = "./config/"
)

var cfg _Config

type _Config struct {
	Database _Database `json:"database"`
	Server   _Server   `json:"server"`
}

type _Server struct {
	Port string `json:"port"`
}

type _Database struct {
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbName"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

// Peek provides secure access to config options
func Peek() _Config {
	return cfg
}

func init() {
	var (
		cfgFile = configPath + "local.json"
		err     error
	)

	file, err := os.Open(cfgFile)

	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(file).Decode(&cfg)
	if err != nil {
		panic(err)
	}
	log.Println("config parsed from", "local.json")
}
