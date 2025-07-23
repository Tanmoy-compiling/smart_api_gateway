package main

import (
	"log"
	"net/http"

	"smart_api_gateway/pkg/common/conf"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment only")
	}
    
	if err := conf.InitDB(); err != nil {
		log.Fatalf("Failed to initialize DB: %v", err)
	}
	defer conf.CloseDB()

	r := gin.Default()

	// Example route just to check app is running
	r.GET("/health", func(c *gin.Context) {
		// Optionally check DB connection health
		if err := conf.DB.Ping(c); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "db unreachable"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	log.Println("Product service running on :5002")
	r.Run(":5002") // Start HTTP server
}
