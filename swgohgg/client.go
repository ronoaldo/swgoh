package swgohgg

import (
	"net/http"
)

type Client struct {
	hc      *http.Client
	profile string
}

func NewClient(profile string) *Client {
	return &Client{
		hc:      http.DefaultClient,
		profile: profile,
	}
}

func (c *Client) UseHTTPClient(hc *http.Client) *Client {
	c.hc = hc
	return c
}

func (c *Client) Profile(profile string) *Client {
	c.profile = profile
	return c
}
