package word

import (
	"fmt"
	"net/http"

	"github.com/nesbitjd/hangle_server/database"
	"github.com/nesbitjd/hangle_server/types"

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

	id := c.Param("id")

	logrus.Debugf("delete: %+v\n", db.Where("id = ?", id).Delete(&types.Word{}))

	resp := fmt.Sprintf("deleted entry %+v", id)
	c.JSON(http.StatusOK, resp)
}
