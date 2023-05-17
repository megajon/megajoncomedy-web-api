package main

import (
	"log"

	g "github.com/heroku/go-getting-started/src/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

// var Database *bun.DB

func main() {

	// port := os.Getenv("PORT")
	port := "3000"

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := g.SetupRouter()
	router.Run(":" + port)
}
