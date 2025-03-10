package storage

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

const dbFile = "./subscribers.db"

// InitDB initializes the database and creates the subscribers table if it doesn't exist
func InitDB() error {
	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS subscribers (chat_id TEXT UNIQUE);`)
	return err
}

// AddSubscriber adds a new subscriber to the database
func AddSubscriber(chatID int64) error {
	db, _ := sql.Open("sqlite", dbFile)
	defer db.Close()

	_, err := db.Exec("INSERT INTO subscribers (chat_id) VALUES (?)", chatID)
	return err
}

// RemoveSubscriber removes a subscriber from the database by chat ID
func RemoveSubscriber(chatID int64) error {
	db, _ := sql.Open("sqlite", dbFile)
	defer db.Close()

	_, err := db.Exec("DELETE FROM subscribers WHERE chat_id = ?", chatID)
	return err
}

// GetSubscribers retrieves all subscriber chat IDs from the database
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
