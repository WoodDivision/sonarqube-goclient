package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Errors struct {
	Msg string `json:"msg"`
}

type Client struct {
	url, user, pass string
	HTTPClient      *http.Client
	QualityGates    *QualityGates
}

func NewClient(url string, user string, pass string) *Client {
	c := &Client{
		url:        url,
		user:       user,
		pass:       pass,
		HTTPClient: &http.Client{},
	}
	c.QualityGates = &QualityGates{client: c}
	return c
}

func (c *Client) GetSonarqubeUrl() string {
	return c.url
}

func (c *Client) SendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(c.user, c.pass)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes *Errors
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Msg)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&v)
	if err != nil {
		return err
	}
	return nil
}
