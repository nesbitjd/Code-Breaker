package word

import (
	"fmt"
	"net/http"

	"github.com/nesbitjd/hangle_server/database"
	"github.com/nesbitjd/hangle_server/types"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Read returns word struct from database
func Read(c *gin.Context) {
	logrus.Info("Reading word struct from database")
	db, err := database.Open()
	if err != nil {
		retErr := fmt.Errorf("unable to open database: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}

	id := c.Param("id")
	readingWord := types.Word{}

	logrus.Debug("Scan table for word struct")
	db.Where("id = ?", id).Find(&readingWord).Scan(&readingWord)

	c.JSON(http.StatusOK, readingWord)
}
