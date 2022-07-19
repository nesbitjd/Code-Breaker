package word

import (
	"fmt"
	"net/http"

	"github.com/nesbitjd/hangle_server/database"
	"github.com/nesbitjd/hangle_server/types"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Update database entry for word
func Update(c *gin.Context) {
	logrus.Info("Updating database entry for word")
	db, err := database.Open()
	if err != nil {
		retErr := fmt.Errorf("unable to open database: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}

	id := c.Param("id")
	word := &types.Word{}

	logrus.Trace("Binding requested id to word struct")
	err = c.Bind(word)
	if err != nil {
		retErr := fmt.Errorf("unable to parse json body: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}

	logrus.Debug("Scan table for database entry and update word struct")
	db.Model(&types.Word{}).Where("id = ?", id).Updates(word)

	resp := fmt.Sprintf("updated entry %+v", word.Word)
	c.JSON(http.StatusCreated, resp)
}
