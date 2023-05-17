package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	u "github.com/heroku/go-getting-started/helpers"
	"github.com/uptrace/bun"
)

func GetRoot(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to the megajon-web api")
}

func GetEmails(c *gin.Context) {
	db := u.Connect()
	emails := make([]u.Email, 0)

	err := db.NewRaw("SELECT id, email FROM ?", bun.Ident("emails")).Scan(c, &emails)
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, emails)
}

func RegisterEmail(c *gin.Context) {

	emailInput := u.CreateEmailObject(c.PostForm("email"))
	if emailInput.Email == "invalid" {
		c.JSON(403, gin.H{"message": "invalid email"})
		return
	}

	db := u.Connect()

	_, err := db.NewInsert().Model(&emailInput).Exec(c)
	if err != nil {
		c.JSON(403, gin.H{"message": "database error"})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
	// u.SendEmail()
}

func DeleteEmail(c *gin.Context) {
	emailInput := u.Email{
		Email: c.PostForm("email"),
	}
	isEmailValid := u.ValidateEmail(emailInput)
	if isEmailValid != nil {
		c.JSON(403, gin.H{"message": "invalid email format"})
		return
	}

	db := u.Connect()
	emails := make([]u.Email, 0)

	err := db.NewRaw("SELECT id, email FROM ?", bun.Ident("emails")).Scan(c, &emails)
	if err != nil {
		c.JSON(404, gin.H{"message": err})
	}

	for _, email := range emails {
		emailToDelete := &u.Email{ID: email.ID, Email: email.Email}
		if email.Email == emailInput.Email {
			_, err := db.NewDelete().Model(emailToDelete).WherePK().Exec(c)
			if err != nil {
				c.JSON(404, gin.H{"message": err})
			}
			c.JSON(200, gin.H{"message": "email deleted"})
			return
		}
	}
	c.JSON(404, gin.H{"message": "no email found"})
}

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
