package api

import (
	"Projects/hangle_server/database"
	"Projects/hangle_server/types"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Update database entry for word
func Update(c *gin.Context) {
	logrus.Info("Updating database entry for word")
	db, err := database.Open()
	if err != nil {
		retErr := fmt.Errorf("unable to open database: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}

	id := c.Param("id")
	hangman := &types.Hangman{}

	logrus.Trace("Binding requested id to hangman type")
	err = c.Bind(hangman)
	if err != nil {
		retErr := fmt.Errorf("unable to parse json body: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}

	logrus.Debug("Scan table for database entry and update word struct")
	hDB := hangman.HangmanToDB()
	db.Model(&types.HangmanDB{}).Where("id = ?", id).Updates(hDB)

	resp := fmt.Sprintf("updated entry %+v", hangman.Word)
	c.JSON(http.StatusCreated, resp)
}
