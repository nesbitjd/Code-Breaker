package user

import (
	"fmt"
	"net/http"

	"github.com/nesbitjd/hangle_server/database"
	"github.com/nesbitjd/hangle_server/pkg/hangle"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Create creates a database entry for the given user
func Create(c *gin.Context) {
	logrus.Info("Creating entry for new user")
	db, err := database.Open("postgres")
	if err != nil {
		retErr := fmt.Errorf("unable to open database: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}

	logrus.Debug("Binding json input to hangman struct")
	user := &hangle.User{}
	err = c.Bind(user)
	if err != nil {
		retErr := fmt.Errorf("unable to parse json body: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}

	logrus.Trace("Create UserDB")
	userDB := db.Create(&user)
	if userDB.Error != nil {
		retErr := fmt.Errorf("unable to create user: %w", userDB.Error)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}

	logrus.Debugf("created: %+v\n", userDB)

	userReturn := hangle.User{}
	logrus.Debug("Scan table for user struct")
	db.Where("username = ?", user.Username).Find(&userReturn).Scan(&userReturn)

	c.JSON(http.StatusCreated, userReturn)
}
