package database

import (
	"Projects/code_breaker/types"
	"fmt"

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
func Setup() error {

	db, err := Open()
	if err != nil {
		return fmt.Errorf("unable to open database: %w", err)
	}

	logrus.Debug("Automigrating")
	err = db.AutoMigrate(&types.HangmanDB{})
	if err != nil {
		return fmt.Errorf("unable to automigrate: %w", err)
	}

	return nil
}

func Open() (*gorm.DB, error) {

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})

	return db, err
}
