package api

import (
	"Projects/code_breaker/database"
	"Projects/code_breaker/types"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Update updates a database entry
func Update(c *gin.Context) {

	logrus.Debug("Opening up database")
	db, err := database.Open()
	if err != nil {
		retErr := fmt.Errorf("unable to open database: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}

	id := c.Param("id")
	hangman := &types.Hangman{}
	err = c.Bind(hangman)
	if err != nil {
		retErr := fmt.Errorf("unable to parse json body: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}

	hDB := hangman.HangmanToDB()
	db.Model(&types.HangmanDB{}).Where("id = ?", id).Updates(hDB)

	resp := fmt.Sprintf("updated entry %+v", hangman.Word)
	c.JSON(http.StatusCreated, resp)
}
