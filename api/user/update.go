package user

import (
	"fmt"
	"net/http"

	"github.com/nesbitjd/hangle_server/database"
	"github.com/nesbitjd/hangle_server/types"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Update database entry for user
func Update(c *gin.Context) {
	logrus.Info("Updating database entry for user")
	db, err := database.Open()
	if err != nil {
		retErr := fmt.Errorf("unable to open database: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}

	id := c.Param("id")
	user := &types.User{}

	logrus.Trace("Binding requested id to hangman type")
	err = c.Bind(user)
	if err != nil {
		retErr := fmt.Errorf("unable to parse json body: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}

	logrus.Debug("Scan table for database entry and update user struct")
	db.Model(&types.User{}).Where("id = ?", id).Updates(user)

	resp := fmt.Sprintf("updated entry %+v", user.Username)
	c.JSON(http.StatusCreated, resp)
}
