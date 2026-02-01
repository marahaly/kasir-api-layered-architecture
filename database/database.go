package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDB(connectionString string) (*sql.DB, error) {
	/** CONNECT TO DATABASE **/
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	/** TEST THE CONNECTION **/
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	/** SET CONNECTION POOL SETTING - RECOMMENDED) **/
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	log.Println("Database connected successfully")
	return db, nil
}