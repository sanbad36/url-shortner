package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sanbad36/url-shortner/api/database"
	"github.com/sanbad36/url-shortner/api/routersetup"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	database.Init()

	router := gin.Default()
	routersetup.SetupRouters(router)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on port %s...", port)
	log.Fatal(router.Run(":" + port))
}
