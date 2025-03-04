package storage

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

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

func AddSubscriber(chatID string) error {
	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO subscribers (chat_id) VALUES (?)", chatID)

	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return fmt.Errorf("the user already exists")
		}
		return err
	}

	log.Printf("Added subscriber with chat_id: %s\n", chatID)
	return nil
}

func RemoveSubscriber(chatID string) error {
	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		return err
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM subscribers WHERE chat_id =?", chatID)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		log.Println("No subscriber found with chat_id:", chatID)
	} else {
		log.Println("Removed subscriber with chat_id:", chatID)
	}

	return nil
}

func GetSubscribers() ([]string, error) {
	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var subscribers []string
	rows, err := db.Query("SELECT chat_id FROM subscribers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var chatID string
		if err := rows.Scan(&chatID); err != nil {
			return nil, err
		}
		subscribers = append(subscribers, chatID)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return subscribers, nil
}
