package swgohgg

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
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

func (c *Client) Get(url string) (*goquery.Document, error) {
	resp, err := c.hc.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("swgohgg: unable to find collection for profile '%s'", c.profile)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("swgohgg: unexpected status code %d", resp.StatusCode)
	}
	return goquery.NewDocumentFromReader(resp.Body)
}

func (c *Client) UseHTTPClient(hc *http.Client) *Client {
	c.hc = hc
	return c
}

func (c *Client) Profile(profile string) *Client {
	c.profile = profile
	return c
}
