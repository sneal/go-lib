package client

import (
	"github.com/sneal/go-lib/v3/resource"
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

func (c *Client) GetApp(guid string) *resource.App {
	return &resource.App{
		GUID:      guid,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Name:      "test app v3 - new and improved",
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
