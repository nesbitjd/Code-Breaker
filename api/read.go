package api

import (
	"Projects/hangle_server/database"
	"Projects/hangle_server/types"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Read returns word struct from database
func Read(c *gin.Context) {
	logrus.Info("Reading word struct from database")
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

	logrus.Debug("Scan table for word struct")
	db.Where("id = ?", id).Find(&readingHangman).Scan(&readingHangman)

	returnHangman := readingHangman.DBtoHangman()

	c.JSON(http.StatusOK, returnHangman)
}
