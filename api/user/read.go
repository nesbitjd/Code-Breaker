package user

import (
	"fmt"
	"net/http"

	"github.com/nesbitjd/hangle_server/database"
	"github.com/nesbitjd/hangle_server/types"

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

// ReadAll returns all user structs from database
func ReadAll(c *gin.Context) {
	logrus.Info("Reading user struct from database")
	db, err := database.Open()
	if err != nil {
		retErr := fmt.Errorf("unable to open database: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}

	readingUser := []types.User{}

	logrus.Debug("Scan table for user struct")
	db.Find(&readingUser).Scan(&readingUser)

	c.JSON(http.StatusOK, readingUser)
}
