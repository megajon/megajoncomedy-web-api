package gin

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	router.Use(cors.Default())
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "static")

	router.GET("/", GetRoot)
	router.GET("/emails", GetEmails)
	router.POST("/register", RegisterEmail)
	router.POST("/delete", DeleteEmail)

	return router
}
