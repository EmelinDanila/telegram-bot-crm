package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/EmelinDanila/telegram-bot-crm/internal/storage"
)

func main() {
	// Load environment variables from.env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading.env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable not set")
	}

	err = storage.InitDB()
	if err != nil {
		log.Fatal("Error initializing database: ", err)
	}
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	log.Println("Server started on port: ", port)
	router.Run(":" + port)
}
