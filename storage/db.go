package storage

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
)

var db *sql.DB

func Initialize() {
	databaseHost := os.Getenv("DB_HOST")
	databasePort := os.Getenv("DB_PORT")
	databaseUser := os.Getenv("DB_USER")
	databasePass := os.Getenv("DB_PASS")
	databaseName := os.Getenv("DB_NAME")

	postgresConnectionURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", databaseUser, databasePass, databaseHost, databasePort, databaseName)
	// var err error
	db, err := sql.Open("postgres", postgresConnectionURL)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	maxOpenConn, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONN"))
	if err != nil {
		panic(err)
	}

	maxIdleConn, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONN"))
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(maxOpenConn)
	db.SetMaxIdleConns(maxIdleConn)

	fmt.Println("Database connected")

}

// returndbinstance returns a pointer to this db connection

func returnDbInstance() *sql.DB {
	return db
}

func ping() {
	err := db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Succesfully connected")
}
