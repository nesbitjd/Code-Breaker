package types

import (
	"encoding/json"
	"fmt"
	"net/http"

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

// NewRecord instantiates a new Record struct with the given inputs
func NewRecord(w Word, u User, f int, g string) Record {
	return Record{
		Word:     w,
		User:     u,
		Failures: f,
		Guesses:  g,
	}
}

// PostResults posts the results of a finshed game to the database
func (c *Client) PostRecord(r Record) (*http.Response, error) {
	postBody, err := json.Marshal(r)
	if err != nil {
		return &http.Response{}, fmt.Errorf("unable to marshal json: %w", err)
	}

	resp, err := c.DoHttp(http.MethodPost, "record", postBody)
	if err != nil {
		return &http.Response{}, fmt.Errorf("unable to do http request: %w", err)
	}

	return resp, nil
}

func (c *Client) DeleteRecord(id string) (*http.Response, error) {
	resp, err := c.DoHttp(http.MethodPost, "record", postBody)
	if err != nil {
		return &http.Response{}, fmt.Errorf("unable to do http request: %w", err)
	}

	return resp, nil
}
