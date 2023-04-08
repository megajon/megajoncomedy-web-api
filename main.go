package main

import (
	"log"

	h "github.com/heroku/go-getting-started/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/uptrace/bun"
)

var Database *bun.DB

func main() {

	// port := os.Getenv("PORT")
	port := "3000"

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Static("/static", "static")

	router.GET("/", h.GetRoot)
	router.GET("/emails", h.GetEmails)
	router.POST("/register", h.RegisterEmail)
	router.POST("/delete", h.DeleteEmail)

	router.Run(":" + port)
}
