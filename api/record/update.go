package record

import (
	"fmt"
	"net/http"

	"github.com/nesbitjd/hangle_server/database"
	"github.com/nesbitjd/hangle_server/types"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Update database entry for record
func Update(c *gin.Context) {
	logrus.Info("Updating database entry for record")
	db, err := database.Open()
	if err != nil {
		retErr := fmt.Errorf("unable to open database: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}

	id := c.Param("id")
	record := &types.Record{}

	logrus.Trace("Binding requested id to record type")
	err = c.Bind(record)
	if err != nil {
		retErr := fmt.Errorf("unable to parse json body: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}

	logrus.Debug("Scan table for database entry and update record struct")
	db.Model(&types.Record{}).Where("id = ?", id).Updates(record)

	resp := fmt.Sprintf("updated entry %+v", record.Word)
	c.JSON(http.StatusCreated, resp)
}
