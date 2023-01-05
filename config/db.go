package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() *sql.DB {
	DB = getDBConnection()
	return DB
}

func getDBConnection() *sql.DB {
	db, err := sql.Open("postgres", Config.DB_SERVICE)
	if err != nil {
		log.Panicf("open conexion error: %v", err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(5)

	return db
}
