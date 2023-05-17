package gin

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Static("/static", "static")

	router.GET("/", GetRoot)
	router.GET("/emails", GetEmails)
	router.POST("/register", RegisterEmail)
	router.POST("/delete", DeleteEmail)

	return router
}
