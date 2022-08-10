package types

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path"
	"strings"

	"gorm.io/gorm"
)

var (
	apiBase = "api/v1"
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

// PostResults posts the results of a finshed game to the database
func (r *Record) PostResults(base_url string) error {
	record_url := path.Join(base_url, apiBase, "record")

	postBody, err := json.Marshal(r)
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

// PostUser posts a new user to the database
func (u *User) PostUser(base_url string) error {
	user_url := path.Join(base_url, apiBase, "user")
	postBody, err := json.Marshal(u)
	if err != nil {
		return err
	}

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, user_url, bytes.NewBuffer(postBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &u)
	if err != nil {
		return err
	}

	return nil
}
