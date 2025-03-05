package storage

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

const dbFile = "./subscribers.db"

func InitDB() error {
	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS subscribers (chat_id TEXT UNIQUE);`)
	return err
}

func AddSubscriber(chatID int64) error {
	db, _ := sql.Open("sqlite", dbFile)
	defer db.Close()

	_, err := db.Exec("INSERT INTO subscribers (chat_id) VALUES (?)", chatID)
	return err
}

func RemoveSubscriber(chatID int64) error {
	db, _ := sql.Open("sqlite", dbFile)
	defer db.Close()

	_, err := db.Exec("DELETE FROM subscribers WHERE chat_id = ?", chatID)
	return err
}

func GetSubscribers() ([]string, error) {
	db, _ := sql.Open("sqlite", dbFile)
	defer db.Close()

	rows, _ := db.Query("SELECT chat_id FROM subscribers")
	defer rows.Close()

	var subscribers []string
	for rows.Next() {
		var chatID string
		rows.Scan(&chatID)
		subscribers = append(subscribers, chatID)
	}
	return subscribers, nil
}
