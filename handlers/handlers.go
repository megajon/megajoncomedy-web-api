package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	m "github.com/heroku/go-getting-started/models"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

func GetRoot(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to the megajon-web api")
}

func GetEmails(c *gin.Context) {
	db := Connect()
	emails := make([]m.Email, 0)

	err := db.NewRaw("SELECT id, email FROM ?", bun.Ident("emails")).Scan(c, &emails)
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, emails)
}

func RegisterEmail(c *gin.Context) {
	newEmail := m.Email{
		Email: c.PostForm("email"),
	}
	isEmailValid := m.ValidateEmail(newEmail)
	if isEmailValid != nil {
		c.JSON(403, gin.H{"message": "invalid email format"})
		return
	}

	db := Connect()
	res, err := db.NewInsert().Model(&newEmail).Exec(c)
	if err != nil {
		c.JSON(403, gin.H{"message": "email already exists"})
		return
	}
	c.JSON(200, gin.H{"message": res})
}

func DeleteEmail(c *gin.Context) {
	emailInput := m.Email{
		Email: c.PostForm("email"),
	}
	isEmailValid := m.ValidateEmail(emailInput)
	if isEmailValid != nil {
		c.JSON(403, gin.H{"message": "invalid email format"})
		return
	}

	db := Connect()
	emails := make([]m.Email, 0)

	err := db.NewRaw("SELECT id, email FROM ?", bun.Ident("emails")).Scan(c, &emails)
	if err != nil {
		c.JSON(404, gin.H{"message": err})
	}

	for _, email := range emails {
		emailToDelete := &m.Email{ID: email.ID, Email: email.Email}
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

func Connect() *bun.DB {
	dsn := "postgres://zxhymrzk:Efra4FYrgAWrjJHJBrdg2LCM2bwuOOvp@castor.db.elephantsql.com/zxhymrzk"

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	return db
}
