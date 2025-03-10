package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/EmelinDanila/telegram-bot-crm/internal/storage"
)

type TelegramUpdate struct {
	Message struct {
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
		Text string `json:"text"`
	} `json:"message"`
}

// SendTelegramMessage sends a message to Telegram
func SendTelegramMessage(chatID int64, message string) {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)

	log.Printf("Отправка сообщения в Telegram: chatID=%v, message=%s", chatID, message) // Добавил лог

	chatIDStr := strconv.FormatInt(chatID, 10)
	body, _ := json.Marshal(map[string]string{
		"chat_id": chatIDStr,
		"text":    message,
	})

	_, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Println("Ошибка отправки сообщения:", err)
	}
}

// Sendtosubscribers sends a message to all subscribers
func SendToSubscribers(message string) {
	subscribers, _ := storage.GetSubscribers()

	for _, chatIDStr := range subscribers {
		chatID, err := strconv.ParseInt(chatIDStr, 10, 64)
		if err != nil {
			log.Println("Ошибка при преобразовании chatID:", err)
			continue
		}

		SendTelegramMessage(chatID, message)
	}
}
