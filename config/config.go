package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var ENV *GeneralConfig = nil

//GeneralConfig structur
type GeneralConfig struct {
	DatabaseDSN string
	RedisDSN    string
	Debug       bool
}

//GetEnvirotment Get from environtment file
func init() {

	if ENV == nil {

		if err := godotenv.Load(); err != nil {
			log.Print("No .env file found")
		}
		host := os.Getenv("DB_HOST")
		if host == "" {
			log.Print("no DB_HOST in envirotment")
		}

		port := os.Getenv("DB_PORT")
		if port == "" {
			log.Print("no DB_PORT in envirotment")
		}

		dbname := os.Getenv("DB_NAME")
		if dbname == "" {
			log.Print("no DB_NAME in envirotment")

		}
		usname := os.Getenv("DB_USERNAME")
		if usname == "" {
			log.Print("no DB_USERNAME in envirotment")

		}
		pwd := os.Getenv("DB_PASSWORD")
		if pwd == "" {
			log.Print("no DB_PASSWORD in envirotment")

		}

		debug := os.Getenv("DEBUG")
		if debug == "" {
			log.Print("no DEBUG in envirotment")
		}

		var conf GeneralConfig
		conf.DatabaseDSN = GenerateDatabaseDSN(host, usname, pwd, port, dbname)
		conf.Debug = strings.ToUpper(debug) == "TRUE"
		ENV = &conf

		if ENV.Debug {
			log.Println("Getting Configuration From Envirotment")
			log.Println(conf)
		}

	}
}

func GenerateDatabaseDSN(host string, user string, password string, port string, dbname string) string {
	constr := fmt.Sprintf("host=%s user=%s password=%s port=%s  dbname=%s sslmode=disable",
		host,
		user,
		password,
		port,
		dbname)

	return constr
}
