package api

import (
	"Projects/code_breaker/database"
	"Projects/code_breaker/types"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Create creates a database entry
func Create(c *gin.Context) {

	logrus.Debug("Setting up database")
	db := database.Setup(c)

	// Migrate the schema
	logrus.Debug("Automigrating")
	err := db.AutoMigrate(&types.HangmanDB{})
	if err != nil {
		retErr := fmt.Errorf("database unable to automigrate: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusInternalServerError, retErr.Error())
		return
	}

	hangman := &types.Hangman{}
	err = c.Bind(hangman)
	if err != nil {
		retErr := fmt.Errorf("unable to parse json body: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}

	hangmanDB := hangman.HangmanToDB()
	logrus.Debugf("Convert hangmanDB: %+v", hangmanDB)

	hangmanDBres := db.Create(&hangmanDB)
	logrus.Debug("Create hangmanDBres")

	fmt.Printf("create: %+v\n", hangmanDBres)

	resp := fmt.Sprintf("created entry %+v", hangman.Word)
	c.JSON(http.StatusCreated, resp)
}
