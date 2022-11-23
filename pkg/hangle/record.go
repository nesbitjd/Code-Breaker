package hangle

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"gorm.io/gorm"
)

var recordPath = "record"

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

// PostResults performs an API request to create the given record
func (c *Client) PostRecord(r Record) (Record, error) {
	postBody, err := json.Marshal(r)
	if err != nil {
		return Record{}, fmt.Errorf("unable to marshal json: %w", err)
	}
	resp, err := c.DoHttp(http.MethodPost, recordPath, postBody)
	if err != nil {
		return Record{}, fmt.Errorf("unable to do http request: %w", err)
	}
	if resp.StatusCode != http.StatusCreated {
		body, err := readRespBody(resp)
		if err != nil {
			return Record{}, fmt.Errorf("error doing request: %s", resp.Status)
		}
		return Record{}, fmt.Errorf("error doing request: %s: %s", resp.Status, body)
	}

	respRecord := &Record{}
	err = readAndUnmarshalRespBody(resp, respRecord)
	if err != nil {
		return Record{}, fmt.Errorf("received invalid response: %w", err)
	}
	return *respRecord, nil
}

// PutRecord performs an API request to update the given record
func (c *Client) PutRecord(r Record, id string) (Record, error) {
	postBody, err := json.Marshal(r)
	if err != nil {
		return Record{}, fmt.Errorf("unable to marshal json: %w", err)
	}
	url, err := url.JoinPath(recordPath, id)
	if err != nil {
		return Record{}, fmt.Errorf("invalid record id %q", id)
	}
	resp, err := c.DoHttp(http.MethodPut, url, postBody)
	if err != nil {
		return Record{}, fmt.Errorf("unable to do http request: %w", err)
	}
	if resp.StatusCode != http.StatusCreated {
		body, err := readRespBody(resp)
		if err != nil {
			return Record{}, fmt.Errorf("error doing request: %s", resp.Status)
		}
		return Record{}, fmt.Errorf("error doing request: %s: %s", resp.Status, body)
	}

	respRecord := &Record{}
	err = readAndUnmarshalRespBody(resp, respRecord)
	if err != nil {
		return Record{}, fmt.Errorf("received invalid response: %w", err)
	}
	return *respRecord, nil
}

// GetRecord performs an API request to retrieve the given record
func (c *Client) GetRecord(id string) (Record, error) {
	url, err := url.JoinPath(recordPath, id)
	if err != nil {
		return Record{}, fmt.Errorf("invalid record id %q", id)
	}
	resp, err := c.DoHttp(http.MethodGet, url, nil)
	if err != nil {
		return Record{}, fmt.Errorf("unable to do request: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		body, err := readRespBody(resp)
		if err != nil {
			return Record{}, fmt.Errorf("error doing request: %s", resp.Status)
		}
		return Record{}, fmt.Errorf("error doing request: %s: %s", resp.Status, body)
	}

	respRecord := &Record{}
	err = readAndUnmarshalRespBody(resp, respRecord)
	if err != nil {
		return Record{}, fmt.Errorf("received invalid response: %w", err)
	}
	return *respRecord, nil
}

// GetAllRecords performs an API request to retrieve all users
func (c *Client) GetAllRecords(id string) ([]Record, error) {
	resp, err := c.DoHttp(http.MethodGet, recordPath, nil)
	if err != nil {
		return []Record{}, fmt.Errorf("unable to do request: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		body, err := readRespBody(resp)
		if err != nil {
			return []Record{}, fmt.Errorf("error doing request: %s", resp.Status)
		}
		return []Record{}, fmt.Errorf("error doing request: %s: %s", resp.Status, body)
	}

	respRecord := []Record{}
	err = readAndUnmarshalRespBody(resp, respRecord)
	if err != nil {
		return []Record{}, fmt.Errorf("received invalid response: %w", err)
	}
	return respRecord, nil
}

// DeleteRecord performs an API request to delete the given record
func (c *Client) DeleteRecord(id string) (Record, error) {
	url, err := url.JoinPath(recordPath, id)
	if err != nil {
		return Record{}, fmt.Errorf("invalid record id %q", id)
	}
	resp, err := c.DoHttp(http.MethodDelete, url, nil)
	if err != nil {
		return Record{}, fmt.Errorf("unable to do request: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		body, err := readRespBody(resp)
		if err != nil {
			return Record{}, fmt.Errorf("error doing request: %s", resp.Status)
		}
		return Record{}, fmt.Errorf("error doing request: %s: %s", resp.Status, body)
	}

	respRecord := &Record{}
	err = readAndUnmarshalRespBody(resp, respRecord)
	if err != nil {
		return Record{}, fmt.Errorf("received invalid response: %w", err)
	}
	return *respRecord, nil
}
