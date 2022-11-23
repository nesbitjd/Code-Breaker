package hangle

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"gorm.io/gorm"
)

var wordPath = "word"

// Word is the representation of the word used in a game
type Word struct {
	gorm.Model
	Word string
}

// NewWord instantiates a new Word struct with a given word string
func NewWord(w string) Word {
	return Word{
		Word: w,
	}
}

// PostWord performs an API request to create the given word
func (c *Client) PostWord(w Word) (Word, error) {
	postBody, err := json.Marshal(w)
	if err != nil {
		return Word{}, fmt.Errorf("unable to marshal json: %w", err)
	}
	resp, err := c.DoHttp(http.MethodPost, wordPath, postBody)
	if err != nil {
		return Word{}, fmt.Errorf("unable to do http request: %w", err)
	}
	if resp.StatusCode != http.StatusCreated {
		return Word{}, fmt.Errorf("error doing request: %s", resp.Status)
	}

	bytes, err := readRespBody(resp)
	if err != nil {
		return Word{}, fmt.Errorf("received invalid response: %w", err)
	}

	respWord := Word{}
	err = json.Unmarshal(bytes, &respWord)
	if err != nil {
		return Word{}, fmt.Errorf("received invalid response: %w", err)
	}
	return respWord, nil
}

// PutWord performs an API request to update the given word
func (c *Client) PutWord(w Word, id string) (Word, error) {
	postBody, err := json.Marshal(w)
	if err != nil {
		return Word{}, fmt.Errorf("unable to marshal json: %w", err)
	}
	url, err := url.JoinPath(wordPath, id)
	if err != nil {
		return Word{}, fmt.Errorf("invalid word id %q", id)
	}
	resp, err := c.DoHttp(http.MethodPut, url, postBody)
	if err != nil {
		return Word{}, fmt.Errorf("unable to do http request: %w", err)
	}
	if resp.StatusCode != http.StatusCreated {
		return Word{}, fmt.Errorf("error doing request: %s", resp.Status)
	}

	bytes, err := readRespBody(resp)
	if err != nil {
		return Word{}, fmt.Errorf("received invalid response: %w", err)
	}

	respWord := Word{}
	err = json.Unmarshal(bytes, &respWord)
	if err != nil {
		return Word{}, fmt.Errorf("received invalid response: %w", err)
	}
	return respWord, nil
}

// GetWord performs an API request to retrieve the given word
func (c *Client) GetWord(id string) (Word, error) {
	url, err := url.JoinPath(wordPath, id)
	if err != nil {
		return Word{}, fmt.Errorf("invalid word id %q", id)
	}
	resp, err := c.DoHttp(http.MethodGet, url, nil)
	if err != nil {
		return Word{}, fmt.Errorf("unable to do request: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return Word{}, fmt.Errorf("error doing request: %s", resp.Status)
	}

	bytes, err := readRespBody(resp)
	if err != nil {
		return Word{}, fmt.Errorf("received invalid response: %w", err)
	}

	respWord := Word{}
	err = json.Unmarshal(bytes, &respWord)
	if err != nil {
		return Word{}, fmt.Errorf("received invalid response: %w", err)
	}
	return respWord, nil
}

// GetAllWords performs an API request to retrieve all Words
func (c *Client) GetAllWords() (Word, error) {
	resp, err := c.DoHttp(http.MethodGet, wordPath, nil)
	if err != nil {
		return Word{}, fmt.Errorf("unable to do request: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return Word{}, fmt.Errorf("error doing request: %s", resp.Status)
	}

	bytes, err := readRespBody(resp)
	if err != nil {
		return Word{}, fmt.Errorf("received invalid response: %w", err)
	}

	respWord := Word{}
	err = json.Unmarshal(bytes, &respWord)
	if err != nil {
		return Word{}, fmt.Errorf("received invalid response: %w", err)
	}
	return respWord, nil
}

// GetLastWord performs an API request to retrieve the last word in the database
func (c *Client) GetLastWord() (Word, error) {
	url, err := url.JoinPath(wordPath, "last")
	if err != nil {
		return Word{}, fmt.Errorf("unable to create path for last word")
	}

	resp, err := c.DoHttp(http.MethodGet, url, nil)
	if err != nil {
		return Word{}, fmt.Errorf("unable to do request: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return Word{}, fmt.Errorf("error doing request: %s", resp.Status)
	}

	bytes, err := readRespBody(resp)
	if err != nil {
		return Word{}, fmt.Errorf("received invalid response: %w", err)
	}

	respWord := Word{}
	err = json.Unmarshal(bytes, &respWord)
	if err != nil {
		return Word{}, fmt.Errorf("received invalid response: %w", err)
	}
	return respWord, nil
}

// DeleteWord performs an API request to delete the given word
func (c *Client) DeleteWord(id string) (Word, error) {
	url, err := url.JoinPath(wordPath, id)
	if err != nil {
		return Word{}, fmt.Errorf("invalid word id %q", id)
	}

	resp, err := c.DoHttp(http.MethodDelete, url, nil)
	if err != nil {
		return Word{}, fmt.Errorf("unable to do request: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return Word{}, fmt.Errorf("error doing request: %s", resp.Status)
	}

	bytes, err := readRespBody(resp)
	if err != nil {
		return Word{}, fmt.Errorf("received invalid response: %w", err)
	}

	respWord := Word{}
	err = json.Unmarshal(bytes, &respWord)
	if err != nil {
		return Word{}, fmt.Errorf("received invalid response: %w", err)
	}
	return respWord, nil
}
