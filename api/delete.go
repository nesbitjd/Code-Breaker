package api

import (
	"Projects/code_breaker/database"
	"Projects/code_breaker/types"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Delete deletes database entry
func Delete(c *gin.Context) {

	db := database.Setup(c)

	// Migrate the schema
	db.AutoMigrate(&types.HangmanDB{})

	word := c.Param("word")

	fmt.Printf("delete: %+v\n", db.Where("Word = ?", word).Delete(&types.HangmanDB{}))

	resp := fmt.Sprintf("deleted entry %+v\n", word)
	c.JSON(http.StatusOK, resp)
}
