package main

import (
	"log"
	"os"

	g "github.com/heroku/go-getting-started/src/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/joho/godotenv"
)

// var Database *bun.DB

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	// port := "4000"

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := g.SetupRouter()
	router.Run(":" + port)
}
