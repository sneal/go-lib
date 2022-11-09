package go_lib

import (
	"time"
)

type Client struct {
	URL string
}

func NewClient(url string) *Client {
	return &Client{
		URL: url,
	}
}

func (c *Client) GetApp(guid string) *App {
	return &App{
		GUID:      guid,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Name:      "test app",
		State:     "RUNNING",
		Links:     nil,
	}
}

func (c *Client) Get(path string) string {
	return "GET " + path
}

func (c *Client) Delete(path string) string {
	return "DELETE " + path
}

func (c *Client) Post(path string) string {
	return "POST " + path
}
