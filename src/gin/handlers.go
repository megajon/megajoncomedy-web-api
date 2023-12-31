package gin

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	s "github.com/heroku/go-getting-started/src"
	d "github.com/heroku/go-getting-started/src/db"
	e "github.com/heroku/go-getting-started/src/email"
	"github.com/uptrace/bun"
)

func GetRoot(c *gin.Context) {
	c.JSON(403, gin.H{"message": "404 page not found"})
}

func RegisterEmail(c *gin.Context) {
	fmt.Printf("email from form: %v", c.PostForm("email"))
	emailInput := e.CreateEmailObject(strings.ToLower(c.PostForm("email")))
	fmt.Printf("email input: %v", emailInput)
	if emailInput.Email == "invalid" {
		fmt.Println("invalid email")
		c.JSON(403, gin.H{"message": "invalid email"})
		return
	}

	db := d.Connect()

	_, err := db.NewInsert().Model(&emailInput).Exec(c)
	if err != nil {
		fmt.Println("database error")
		c.JSON(403, gin.H{"message": "database error"})
		db.Close()
		return
	}
	c.JSON(200, gin.H{"message": "success"})
	db.Close()
	e.SendNewSubscriberEmail(emailInput.Email)
	e.SendWelcomeEmail(emailInput.Email)
}

func DeleteEmail(c *gin.Context) {
	emailInput := s.Email{
		Email: strings.ToLower(c.PostForm("email")),
	}
	fmt.Printf("email input: %v", emailInput)
	isEmailValid := e.ValidateEmail(emailInput)
	if isEmailValid != nil {
		c.JSON(403, gin.H{"message": "invalid email"})
		return
	}

	db := d.Connect()
	emails := make([]s.Email, 0)

	err := db.NewRaw("SELECT id, email FROM ?", bun.Ident("emails")).Scan(c, &emails)
	if err != nil {
		c.JSON(404, gin.H{"message": err})
		db.Close()
	}

	for _, email := range emails {
		emailToDelete := &s.Email{ID: email.ID, Email: email.Email}
		if email.Email == emailInput.Email {
			_, err := db.NewDelete().Model(emailToDelete).WherePK().Exec(c)
			if err != nil {
				c.JSON(404, gin.H{"message": err})
				db.Close()
			}
			c.JSON(200, gin.H{"message": "email deleted"})
			db.Close()
			e.UnsubscribeEmail(emailInput.Email)
			return
		}
	}
	c.JSON(200, gin.H{"message": "no email found"})
	db.Close()
}
