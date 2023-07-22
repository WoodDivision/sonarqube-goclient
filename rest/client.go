package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/go-querystring/query"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Errors struct {
	Msg string `json:"msg"`
}

type Client struct {
	url, user, pass string
	baseUrl         *url.URL
	HTTPClient      *http.Client
	QualityGates    *QualityGates
}

func NewClient(sonarUrl string, user string, pass string) *Client {
	pars, err := url.Parse(sonarUrl)
	if err != nil {
		log.Printf("Wrong url for SonarQube")
		return nil
	}
	c := &Client{
		baseUrl:    pars,
		user:       user,
		pass:       pass,
		HTTPClient: &http.Client{},
	}
	c.QualityGates = &QualityGates{client: c}
	return c
}

func (c *Client) SendRequest(method string, u *url.URL, v interface{}) error {
	req := &http.Request{
		Method:     method,
		URL:        u,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header),
		Host:       u.Host,
	}
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
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Unknow Error in Body.Close")
			return
		}
	}(res.Body)
	err = json.NewDecoder(res.Body).Decode(&v)
	if err != nil {
		return err
	}
	log.Printf("%#v", v)
	return nil
}

func (c *Client) SetUrlOpt(path string, opt interface{}) *url.URL {
	u := *c.baseUrl
	u.Opaque = u.Path + path

	if opt != nil {
		q, err := query.Values(opt)
		if err != nil {
			log.Printf("wrong req")
			return nil
		}
		u.RawQuery = q.Encode()
	}
	return &u
}
