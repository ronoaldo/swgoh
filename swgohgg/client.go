package swgohgg

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/hashicorp/golang-lru"
)

var globalCache, _ = lru.New(1024)

type Client struct {
	hc       *http.Client
	profile  string
	useCache bool
}

func NewClient(profile string) *Client {
	return &Client{
		hc:      http.DefaultClient,
		profile: profile,
	}
}

func (c *Client) Get(url string) (*goquery.Document, error) {
	b, ok := globalCache.Get(url)
	if ok && c.useCache {
		// Already on cache, use it!
		return goquery.NewDocumentFromReader(bytes.NewBuffer(b.([]byte)))
	}
	// Not in cache, fetch from remote site
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
	data, err := ioutil.ReadAll(resp.Body)
	if c.useCache {
		globalCache.Add(url, data)
	}
	return goquery.NewDocumentFromReader(bytes.NewBuffer(data))
}

func (c *Client) UseHTTPClient(hc *http.Client) *Client {
	c.hc = hc
	return c
}

func (c *Client) Profile(profile string) *Client {
	c.profile = profile
	return c
}

func (c *Client) UseCache(useCache bool) *Client {
	c.useCache = useCache
	return c
}

func FlushCache() {
	globalCache.Purge()
}
