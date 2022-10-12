package types

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"gorm.io/gorm"
)

// Word is the representation of the word used in a game
type Word struct {
	gorm.Model
	Word string
}

// NewWord instantiates a new Word struct with a given word string
func NewWord(w string) *Word {
	return &Word{
		Word: w,
	}
}

// PostUser posts a new word to the database
func (w *Word) PostUser(base_url string) error {
	word_url, err := url.JoinPath(base_url, apiBase, "word")
	if err != nil {
		return fmt.Errorf("unable to parse url: %w", err)
	}

	postBody, err := json.Marshal(w)
	if err != nil {
		return fmt.Errorf("unable to marshal json: %w", err)
	}

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, word_url, bytes.NewBuffer(postBody))
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
