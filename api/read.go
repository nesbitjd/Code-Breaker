package api

import (
	"Projects/code_breaker/database"
	"Projects/code_breaker/types"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Read reads
func Read(c *gin.Context) {

	logrus.Debug("Opening up database")
	db, err := database.Open()
	if err != nil {
		retErr := fmt.Errorf("unable to open database: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}

	id := c.Param("id")
	readingHangman := types.HangmanDB{}

	db.Where("id = ?", id).Find(&readingHangman).Scan(&readingHangman)

	resp := fmt.Sprintf("%+v for ID: %+v", readingHangman, id)

	c.JSON(http.StatusOK, resp)
}
