package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	s "github.com/heroku/go-getting-started/src"
	d "github.com/heroku/go-getting-started/src/db"
	e "github.com/heroku/go-getting-started/src/email"
	"github.com/uptrace/bun"
)

func GetRoot(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to the megajon-web api")
}

func GetEmails(c *gin.Context) {
	db := d.Connect()
	emails := make([]s.Email, 0)

	err := db.NewRaw("SELECT id, email FROM ?", bun.Ident("emails")).Scan(c, &emails)
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, emails)
}

func RegisterEmail(c *gin.Context) {

	emailInput := e.CreateEmailObject(c.PostForm("email"))
	if emailInput.Email == "invalid" {
		c.JSON(403, gin.H{"message": "invalid email"})
		return
	}

	db := d.Connect()

	_, err := db.NewInsert().Model(&emailInput).Exec(c)
	if err != nil {
		c.JSON(403, gin.H{"message": "database error"})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
	// u.SendEmail()
}

func DeleteEmail(c *gin.Context) {
	emailInput := s.Email{
		Email: c.PostForm("email"),
	}
	isEmailValid := e.ValidateEmail(emailInput)
	if isEmailValid != nil {
		c.JSON(403, gin.H{"message": "invalid email format"})
		return
	}

	db := d.Connect()
	emails := make([]s.Email, 0)

	err := db.NewRaw("SELECT id, email FROM ?", bun.Ident("emails")).Scan(c, &emails)
	if err != nil {
		c.JSON(404, gin.H{"message": err})
	}

	for _, email := range emails {
		emailToDelete := &s.Email{ID: email.ID, Email: email.Email}
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
