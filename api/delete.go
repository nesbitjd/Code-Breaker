package api

import (
	"Projects/hangle_server/database"
	"Projects/hangle_server/types"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Delete deletes entry for the given word
func Delete(c *gin.Context) {
	logrus.Info("Deleting entry for the given word")
	db, err := database.Open()
	if err != nil {
		retErr := fmt.Errorf("unable to open database: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}

	word := c.Param("word")

	logrus.Debugf("delete: %+v\n", db.Where("Word = ?", word).Delete(&types.HangmanDB{}))

	resp := fmt.Sprintf("deleted entry %+v", word)
	c.JSON(http.StatusOK, resp)
}
