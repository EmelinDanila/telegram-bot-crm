package main

import (
	"log"
	"os"

	"github.com/EmelinDanila/telegram-bot-crm/internal/handlers"
	"github.com/EmelinDanila/telegram-bot-crm/internal/storage"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// We load the variables of the environment
	err := godotenv.Load()
	if err != nil {
		log.Println("Ошибка загрузки .env файла")
	}

	// Initialize the database
	err = storage.InitDB()
	if err != nil {
		log.Fatal("Ошибка инициализации БД:", err)
	}

	router := gin.Default()

	// EndPome for test webhuki (AMOCRM emulation)
	router.POST("/send", handlers.TestWebhookHandler)

	//Endpoint for processing Telegram Bota commands
	router.POST("/telegram", handlers.TelegramHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Сервер запущен на порту:", port)
	router.Run(":" + port)
}
