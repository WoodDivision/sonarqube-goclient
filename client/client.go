package client

import (
	"net/http"
)

type Client struct {
	url, user, pass string
	HTTPClient      *http.Client
}

func NewClient(url string, user string, pass string) *Client {
	return &Client{
		url:        url,
		user:       user,
		pass:       pass,
		HTTPClient: &http.Client{},
	}
}

func (c *Client) GetSonarqubeUser() string {
	return c.user
}
func (c *Client) GetSonarqubePass() string {
	return c.pass
}
func (c *Client) GetSonarqubeUrl() string {
	return c.url
}
