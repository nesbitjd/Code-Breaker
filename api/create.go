package api

import (
	"Projects/hangle_server/database"
	"Projects/hangle_server/types"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// CreateWord creates a database entry for the given word
func CreateWord(c *gin.Context) {
	logrus.Info("Creating entry for new word")
	db, err := database.Open()
	if err != nil {
		retErr := fmt.Errorf("unable to open database: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}

	logrus.Debug("Binding json input to hangman struct")
	hangman := &types.Hangman{}
	err = c.Bind(hangman)
	if err != nil {
		retErr := fmt.Errorf("unable to parse json body: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}

	logrus.Trace("Convert to hangmanDB")
	hangmanDB := hangman.HangmanToDB()

	logrus.Trace("Create hangmanDBres")
	hangmanDBres := db.Create(&hangmanDB)

	logrus.Debugf("created: %+v\n", hangmanDBres)

	resp := fmt.Sprintf("created entry %+v", hangman.Word)
	c.JSON(http.StatusCreated, resp)
}
