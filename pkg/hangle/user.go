package hangle

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"gorm.io/gorm"
)

var userPath = "user"

// User is the representation of a player credentials
type User struct {
	gorm.Model
	Username string `gorm:"type:text;unique"`
}

// NewUser instantiates a new User struct with the given username
func NewUser(u string) User {
	return User{
		Username: u,
	}
}

// Post performs an API request to create the given user
func (c *Client) PostUser(u User) (User, error) {
	postBody, err := json.Marshal(u)
	if err != nil {
		return User{}, fmt.Errorf("unable to marshal json: %w", err)
	}

	resp, err := c.DoHttp(http.MethodPost, userPath, postBody)
	if err != nil {
		return User{}, fmt.Errorf("unable to do http request: %w", err)
	}
	if resp.StatusCode != http.StatusCreated {
		body, err := readRespBody(resp)
		if err != nil {
			return User{}, fmt.Errorf("error doing request: %s", resp.Status)
		}
		return User{}, fmt.Errorf("error doing request: %s: %s", resp.Status, body)
	}

	respUser := &User{}
	err = readAndUnmarshalRespBody(resp, respUser)
	if err != nil {
		return User{}, fmt.Errorf("received invalid response: %w", err)
	}
	return *respUser, nil
}

// PutUser performs an API request to update the given user
func (c *Client) PutUser(u User, id string) (User, error) {
	postBody, err := json.Marshal(u)
	if err != nil {
		return User{}, fmt.Errorf("unable to marshal json: %w", err)
	}

	url, err := url.JoinPath(userPath, id)
	if err != nil {
		return User{}, fmt.Errorf("invalid user id %q", id)
	}

	resp, err := c.DoHttp(http.MethodPut, url, postBody)
	if err != nil {
		return User{}, fmt.Errorf("unable to do http request: %w", err)
	}
	if resp.StatusCode != http.StatusCreated {
		body, err := readRespBody(resp)
		if err != nil {
			return User{}, fmt.Errorf("error doing request: %s", resp.Status)
		}
		return User{}, fmt.Errorf("error doing request: %s: %s", resp.Status, body)
	}

	respUser := &User{}
	err = readAndUnmarshalRespBody(resp, respUser)
	if err != nil {
		return User{}, fmt.Errorf("received invalid response: %w", err)
	}
	return *respUser, nil
}

// GetAllUsers performs an API request to retrieve all users
func (c *Client) GetAllUsers() ([]User, error) {
	resp, err := c.DoHttp(http.MethodGet, userPath, nil)
	if err != nil {
		return []User{}, fmt.Errorf("unable to do request: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		body, err := readRespBody(resp)
		if err != nil {
			return []User{}, fmt.Errorf("error doing request: %s", resp.Status)
		}
		return []User{}, fmt.Errorf("error doing request: %s: %s", resp.Status, body)
	}

	respUser := []User{}
	err = readAndUnmarshalRespBody(resp, &respUser)
	if err != nil {
		return []User{}, fmt.Errorf("received invalid response: %w", err)
	}

	return respUser, nil
}

// GetUser performs an API request to retrieve the given user
func (c *Client) GetUser(id string) (User, error) {
	url, err := url.JoinPath(userPath, id)
	if err != nil {
		return User{}, fmt.Errorf("invalid user id %q", id)
	}
	resp, err := c.DoHttp(http.MethodGet, url, nil)
	if err != nil {
		return User{}, fmt.Errorf("unable to do request: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		body, err := readRespBody(resp)
		if err != nil {
			return User{}, fmt.Errorf("error doing request: %s", resp.Status)
		}
		return User{}, fmt.Errorf("error doing request: %s: %s", resp.Status, body)
	}

	respUser := &User{}
	err = readAndUnmarshalRespBody(resp, respUser)
	if err != nil {
		return User{}, fmt.Errorf("received invalid response: %w", err)
	}
	return *respUser, nil
}

// DeleteUser performs an API request to delete the given User
func (c *Client) DeleteUser(id string) (User, error) {
	url, err := url.JoinPath(userPath, id)
	if err != nil {
		return User{}, fmt.Errorf("invalid user id %q", id)
	}
	resp, err := c.DoHttp(http.MethodDelete, url, nil)
	if err != nil {
		return User{}, fmt.Errorf("unable to do request: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		body, err := readRespBody(resp)
		if err != nil {
			return User{}, fmt.Errorf("error doing request: %s", resp.Status)
		}
		return User{}, fmt.Errorf("error doing request: %s: %s", resp.Status, body)
	}

	respUser := &User{}
	err = readAndUnmarshalRespBody(resp, respUser)
	if err != nil {
		return User{}, fmt.Errorf("received invalid response: %w", err)
	}
	return *respUser, nil
}
