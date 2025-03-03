package storage

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

const dbFile = "subscribes.db"

func InitDB() error {
	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS subscribers (
		chat_id TEXT UNIQUE
	);`)

	if err != nil {
		log.Println("Error creating table subscribers", err)
		return err
	}

	log.Println("Created table subscribers")
	return nil
}
