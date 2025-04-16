package main

import (
	"log"
	"os"
	"subly/data"
	"sync"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
)

const PORT = "80"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading Environment file")
	}
	//DataBase
	db := initDB()
	// db.Ping()

	// Session
	session := initSession()

	// Loggers
	infoLog := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR:\t", log.Ldate|log.Ltime)

	// Waitgroups
	wg := sync.WaitGroup{}

	// application config
	app := Config{
		Session:  session,
		DB:       db,
		Wait:     &wg,
		ErrorLog: errorLog,
		InfoLog:  infoLog,
		Models:   data.New(db),
	}

	app.Mailer = app.createMail()
	go app.listenForMail()

	// listen for web connection
	app.server()
}
