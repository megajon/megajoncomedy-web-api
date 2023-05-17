package main

import (
	"log"

	h "github.com/heroku/go-getting-started/handlers"

	_ "github.com/heroku/x/hmetrics/onload"
)

// var Database *bun.DB

func main() {

	// port := os.Getenv("PORT")
	port := "3000"

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := h.SetupRouter()
	router.Run(":" + port)
}
