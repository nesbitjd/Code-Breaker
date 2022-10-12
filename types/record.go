package types

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

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
func (r *Record) PostResults(base_url string) error {
	record_url, err := url.JoinPath(base_url, apiBase, "record")
	if err != nil {
		return fmt.Errorf("unable to parse url: %w", err)
	}

	postBody, err := json.Marshal(r)
	if err != nil {
		return fmt.Errorf("unable to marshal json: %w", err)
	}

	doHttp(http.MethodPost, record_url, postBody)

	return nil
}

func doHttp(method string, url string, data []byte) (*http.Response, error) {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return &http.Response{}, fmt.Errorf("unable to wrap NewRequest: %w", err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return &http.Response{}, fmt.Errorf("error from sending request: %w", err)

	}

	return resp, nil
}
