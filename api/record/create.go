package record

import (
	"fmt"
	"net/http"

	"github.com/nesbitjd/hangle_server/database"
	"github.com/nesbitjd/hangle_server/pkg/hangle"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Create creates a database entry for the given record
func Create(c *gin.Context) {
	logrus.Info("Creating entry for new record")
	db, err := database.Open("postgres")
	if err != nil {
		retErr := fmt.Errorf("unable to open database: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}

	logrus.Debug("Binding json input to record struct")
	record := &hangle.Record{}
	err = c.Bind(record)
	if err != nil {
		retErr := fmt.Errorf("unable to parse json body: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}

	logrus.Trace("Create recordDB")
	recordDB := db.Create(&record)

	logrus.Debugf("created: %+v\n", recordDB)

	recordReturn := hangle.Record{}
	logrus.Debug("Scan table for record struct")
	db.Where("username = ?", record.User.Username).Where("word = ?", record.Word.Word).Find(&recordReturn).Scan(&recordReturn)

	c.JSON(http.StatusCreated, recordReturn)
}
