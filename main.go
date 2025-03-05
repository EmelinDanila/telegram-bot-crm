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
	// Загружаем переменные окружения
	err := godotenv.Load()
	if err != nil {
		log.Println("Ошибка загрузки .env файла")
	}

	// Инициализируем базу данных
	err = storage.InitDB()
	if err != nil {
		log.Fatal("Ошибка инициализации БД:", err)
	}

	router := gin.Default()

	// Эндпоинт для тестовых вебхуков (эмуляция AmoCRM)
	router.POST("/send", handlers.TestWebhookHandler)

	// Эндпоинт для обработки команд Telegram-бота
	router.POST("/telegram", handlers.TelegramHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Сервер запущен на порту:", port)
	router.Run(":" + port)
}
