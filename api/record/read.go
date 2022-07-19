package record

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/nesbitjd/hangle_server/database"
	"github.com/nesbitjd/hangle_server/types"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Read returns record struct from database
func Read(c *gin.Context) {
	logrus.Info("Reading record struct from database")
	db, err := database.Open()
	if err != nil {
		retErr := fmt.Errorf("unable to open database: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		retErr := fmt.Errorf("unable to convert id to string: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}
	readingRecord := types.Record{}

	logrus.Debug("Scan table for record struct")
	// db.Where("id = ?", id).Find(&readingRecord).Scan(&readingRecord).Preload("Word").Preload("User")

	db.Preload("Word").Preload("User").Where("id = ?", id).Find(&readingRecord)

	c.JSON(http.StatusOK, readingRecord)
}
