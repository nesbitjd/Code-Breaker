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

// User is the representation of a player credentials
type User struct {
	gorm.Model
	Username string `gorm:"type:text;unique"`
}

// NewUser instantiates a new User struct with the given username
func NewUser(u string) *User {
	return &User{
		Username: u,
	}
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
	if resp.StatusCode != 201 {
		error_message := ""
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("unable to read body: %w", err)
		}

		err = json.Unmarshal(body, &error_message)
		if err != nil {
			return fmt.Errorf("unable to unmarshal json: %w", err)
		}

		return fmt.Errorf("server returned: %s", error_message)
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
