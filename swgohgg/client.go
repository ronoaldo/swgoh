package swgohgg

import (
	"net/http"
)

type Client struct {
	hc *http.Client
}

func NewClient() *Client {
	return &Client{
		hc: http.DefaultClient,
	}
}

func (c *Client) UseHTTPClient(hc *http.Client) *Client {
	c.hc = hc
	return c
}
