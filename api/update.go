package api

import (
	"Projects/code_breaker/database"
	"Projects/code_breaker/types"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Update updates a database entry
func Update(c *gin.Context) {

	db := database.Setup(c)

	// Migrate the schema
	db.AutoMigrate(&types.HangmanDB{})

	word := c.Param("word")
	hangman := &types.Hangman{}
	err := c.Bind(hangman)
	if err != nil {
		retErr := fmt.Errorf("unable to parse json body: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}

	fmt.Printf("update: %+v\n", db.Model(&hangman).Where("Word = ?", word).Updates(types.HangmanDB{Word: hangman.Word, Failures: hangman.Failures}))

	resp := fmt.Sprintf("updated entry %+v", hangman.Word)
	c.JSON(http.StatusCreated, resp)
}
