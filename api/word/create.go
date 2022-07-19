package word

import (
	"fmt"
	"net/http"

	"github.com/nesbitjd/hangle_server/database"
	"github.com/nesbitjd/hangle_server/types"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Create creates a database entry for the given word
func Create(c *gin.Context) {
	logrus.Info("Creating entry for new word")
	db, err := database.Open()
	if err != nil {
		retErr := fmt.Errorf("unable to open database: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}

	logrus.Debug("Binding json input to word struct")
	word := &types.Word{}
	err = c.Bind(word)
	if err != nil {
		retErr := fmt.Errorf("unable to parse json body: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}

	logrus.Trace("Create wordDB")
	wordDB := db.Create(&word)

	logrus.Debugf("created: %+v\n", wordDB)

	resp := fmt.Sprintf("created entry %+v", word.Word)
	c.JSON(http.StatusCreated, resp)
}
