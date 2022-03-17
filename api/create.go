package api

import (
	"Projects/code_breaker/database"
	"Projects/code_breaker/types"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create creates a database entry
func Create(c *gin.Context) {

	db := database.Setup(c)

	// Migrate the schema
	db.AutoMigrate(&types.HangmanDB{})

	hangman := &types.Hangman{}
	err := c.Bind(hangman)
	if err != nil {
		retErr := fmt.Errorf("unable to parse json body: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}

	hangmanDB := hangman.HangmanToDB()

	fmt.Printf("create: %+v\n", db.Create(hangmanDB))

	resp := fmt.Sprintf("created entry %+v", hangmanDB.Word)
	c.JSON(http.StatusCreated, resp)
}
