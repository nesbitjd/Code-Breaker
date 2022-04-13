package api

import (
	"Projects/code_breaker/database"
	"Projects/code_breaker/types"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Delete deletes database entry
func Delete(c *gin.Context) {

	logrus.Debug("Opening up database")
	db, err := database.Open()
	if err != nil {
		retErr := fmt.Errorf("unable to open database: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}

	word := c.Param("word")

	fmt.Printf("delete: %+v\n", db.Where("Word = ?", word).Delete(&types.HangmanDB{}))

	resp := fmt.Sprintf("deleted entry %+v", word)
	c.JSON(http.StatusOK, resp)
}
