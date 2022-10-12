package database

import (
	"fmt"

	"github.com/nesbitjd/hangle_server/types"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	port     = 5432
	user     = "postgres"
	password = "foobar"
	dbname   = "postgres"
)

// Setup sets up the database
func Setup() error {
	db, err := Open("postgres")
	db.Logger.LogMode(logger.Info)
	if err != nil {
		return fmt.Errorf("unable to open database: %w", err)
	}

	logrus.Debug("Automigrating database")
	err = db.AutoMigrate(&types.User{}, &types.Word{}, &types.Record{})
	if err != nil {
		return fmt.Errorf("unable to automigrate: %w", err)
	}

	return nil
}

func Open(host string) (*gorm.DB, error) {
	logrus.Info("Opening connection to database")
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})

	return db, err
}
