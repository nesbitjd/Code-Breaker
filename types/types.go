package types

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"gorm.io/gorm"
)

// Record is the representation of a finished game state
type Record struct {
	gorm.Model
	User     User
	Word     Word
	WordID   int
	UserID   int
	Failures int
	Guesses  string
}

// Word is the representation of the word used in a game
type Word struct {
	gorm.Model
	Word string
}

// User is the representation of a player credentials
type User struct {
	gorm.Model
	Username string `gorm:"type:text"`
}

// Hangman is used to track the game state
type Hangman struct {
	Word     string
	Failures int
	Guesses  Guesses
}

// Guesses tracks the guessed letters in the game
type Guesses []string

// ToString converts Guesses into a string
func (g *Guesses) ToString() string {

	return strings.Join(*g, ", ")
}

// GuessesFromString converts a string into Guesses
func GuessesFromString(g string) Guesses {

	return strings.Split(g, ", ")
}

// UploadResults uploads the results of a finshed game to the database
func (game *Record) UploadResults(base_url string, failures int, guesses []string) error {

	game.Failures = failures
	game.Guesses = strings.Join(guesses, ",")
	game.UserID = 1

	record_url := fmt.Sprintf("%s/record", base_url)

	postBody, err := json.Marshal(game)
	if err != nil {
		return err
	}

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, record_url, bytes.NewBuffer(postBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	_, err = client.Do(req)
	if err != nil {
		return err
	}

	return nil
}
