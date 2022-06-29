package user

import (
	"Projects/hangle_server/database"
	"Projects/hangle_server/types"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Read returns user struct from database
func Read(c *gin.Context) {
	logrus.Info("Reading user struct from database")
	db, err := database.Open()
	if err != nil {
		retErr := fmt.Errorf("unable to open database: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}

	id := c.Param("id")
	readingUser := types.User{}

	logrus.Debug("Scan table for user struct")
	db.Where("id = ?", id).Find(&readingUser).Scan(&readingUser)

	c.JSON(http.StatusOK, readingUser)
}
