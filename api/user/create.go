package user

import (
	"Projects/hangle_server/database"
	"Projects/hangle_server/types"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Create creates a database entry for the given user
func Create(c *gin.Context) {
	logrus.Info("Creating entry for new user")
	db, err := database.Open()
	if err != nil {
		retErr := fmt.Errorf("unable to open database: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}

	logrus.Debug("Binding json input to hangman struct")
	user := &types.User{}
	err = c.Bind(user)
	if err != nil {
		retErr := fmt.Errorf("unable to parse json body: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}

	logrus.Trace("Create UserDB")
	userDB := db.Create(&user)

	logrus.Debugf("created: %+v\n", userDB)

	resp := fmt.Sprintf("created entry %+v", user.Username)
	c.JSON(http.StatusCreated, resp)
}
