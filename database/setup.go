package database

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	password = "foobar"
	dbname   = "postgres"
)

// Setup sets up the database
func Setup(c *gin.Context) *gorm.DB {

	logrus.SetLevel(logrus.DebugLevel)
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
