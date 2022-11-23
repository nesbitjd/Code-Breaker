package types

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
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
	record_url, err := url.JoinPath(base_url, apiBase, "record")
	if err != nil {
		return fmt.Errorf("unable to parse url: %w", err)
	}

	postBody, err := json.Marshal(r)
	if err != nil {
		return fmt.Errorf("unable to marshal json: %w", err)
	}

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, record_url, bytes.NewBuffer(postBody))
	if err != nil {
		return fmt.Errorf("unable to wrap NewRequest: %w", err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	_, err = client.Do(req)
	if err != nil {
		return fmt.Errorf("error from sending request: %w", err)
	}

	return nil
}

// PostUser posts a new user to the database
func (u *User) PostUser(base_url string) error {
	user_url, err := url.JoinPath(base_url, apiBase, "user")
	if err != nil {
		return fmt.Errorf("unable to parse url: %w", err)
	}

	postBody, err := json.Marshal(u)
	if err != nil {
		return fmt.Errorf("unable to marshal json: %w", err)
	}

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, user_url, bytes.NewBuffer(postBody))
	if err != nil {
		return fmt.Errorf("unable to wrap NewRequest: %w", err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error from sending request: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("unable to read body: %w", err)
	}

	err = json.Unmarshal(body, &u)
	if err != nil {
		return fmt.Errorf("unable to unmarshal json: %w", err)
	}

	return nil
}

// GetLastWord gets the most recently added word
func GetLastWord(base_url string) (Word, error) {
	word_url, err := url.JoinPath(base_url, apiBase, "word/last")
	if err != nil {
		return Word{}, fmt.Errorf("unable to parse url: %w", err)
	}
	resp, err := http.Get(word_url)
	if err != nil {
		return Word{}, fmt.Errorf("unable to get last word: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Word{}, fmt.Errorf("unable to read body: %w", err)
	}

	word := Word{}
	err = json.Unmarshal(body, &word)
	if err != nil {
		return Word{}, fmt.Errorf("unable to unmarshal json: %w", err)
	}

	return word, nil
}

// GetAllUsers returns a slice of all users
func GetAllUsers(base_url string) ([]User, error) {
	user_url, err := url.JoinPath(base_url, apiBase, "user")
	if err != nil {
		return []User{}, fmt.Errorf("unable to parse url: %w", err)
	}

	resp, err := http.Get(user_url)
	if err != nil {
		return []User{}, fmt.Errorf("unable to lookup user: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []User{}, fmt.Errorf("unable to read user body: %w", err)
	}

	users := []User{}
	err = json.Unmarshal(body, &users)
	if err != nil {
		return []User{}, fmt.Errorf("unable to unmarshal user body: %w", err)
	}

	return users, nil
}
