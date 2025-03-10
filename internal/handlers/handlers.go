package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/EmelinDanila/telegram-bot-crm/internal/services"
	"github.com/EmelinDanila/telegram-bot-crm/internal/storage"
	"github.com/gin-gonic/gin"
)

// TelegramHandler processes the commands of the bot
func TelegramHandler(c *gin.Context) {
	var update services.TelegramUpdate
	if err := c.ShouldBindJSON(&update); err != nil {
		log.Println("Ошибка обработки Telegram-данных:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	chatID := update.Message.Chat.ID
	text := update.Message.Text

	switch text {
	case "/subscribe":
		err := storage.AddSubscriber(chatID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to subscribe"})
			return
		}
		services.SendTelegramMessage(chatID, "Вы подписались на уведомления!")

	case "/unsubscribe":
		err := storage.RemoveSubscriber(chatID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to unsubscribe"})
			return
		}
		services.SendTelegramMessage(chatID, "Вы отписались от уведомлений!")
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// Testwebhookhandler Emulor Vedhuk from AMOCRM
func TestWebhookHandler(c *gin.Context) {
	var payload map[string]interface{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		log.Println("Ошибка парсинга вебхука:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	message, _ := json.Marshal(payload)
	log.Println("Получено уведомление:", string(message))

	services.SendToSubscribers(string(message))

	c.JSON(http.StatusOK, gin.H{"status": "notification sent"})
}
