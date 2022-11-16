package types

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
)

var (
	apiBase = "api/v1"
)

type Config struct {
	ServerAddr string
}

func NewConfig(ServerAddr string) Config {
	return Config{
		ServerAddr: ServerAddr,
	}
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	Config Config
	Http   HttpClient
}

func NewClient(config Config, http HttpClient) Client {
	return Client{
		Config: config,
		Http:   http,
	}
}

func (c *Client) DoHttp(method string, endpoint string, data []byte) (*http.Response, error) {
	url, err := url.JoinPath(c.Config.ServerAddr, apiBase, endpoint)
	if err != nil {
		return &http.Response{}, fmt.Errorf("unable to parse url: %w", err)
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return &http.Response{}, fmt.Errorf("unable to wrap NewRequest: %w", err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := c.Http.Do(req)
	if err != nil {
		return &http.Response{}, fmt.Errorf("error from sending request: %w", err)

	}

	return resp, nil
}
