package swgohgg

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

// Client implements methods to interact with the https://swgoh.gg/ website.
type Client struct {
	hc         *http.Client
	jar        *cookiejar.Jar
	profile    string
	authorized bool
}

// NewClient initializes a new instance of the client, tied to the specified user profile.
func NewClient(profile string) *Client {
	jar, err := cookiejar.New(nil)
	if err != nil {
		// Should never happen.
		panic(err)
	}

	c := &Client{
		hc:      http.DefaultClient,
		profile: profile,
	}
	c.hc.Jar = jar
	return c
}

// Get retrieves the provided URL and returns a parsed goquery.Document.
func (c *Client) Get(url string) (*goquery.Document, error) {
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
	return goquery.NewDocumentFromReader(bytes.NewBuffer(data))
}

// UseHTTPClient allows one to overwrite the default HTTP Client.
// The Client.Jar is replaced before next use.
func (c *Client) UseHTTPClient(hc *http.Client) *Client {
	c.hc = hc
	c.hc.Jar = c.jar
	return c
}

// Profile sets the client profile to a new value.
func (c *Client) Profile(profile string) *Client {
	c.profile = profile
	return c
}

// Login authorizes the bot client using the provided username and password.
func (c *Client) Login(username, password string) (err error) {
	resp, err := c.hc.Get("https://swgoh.gg/accounts/login/")
	if err != nil {
		return err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("swgoh.gg: unexpected status code %d: %v", resp.StatusCode, resp.Status)
	}
	loginPage, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return err
	}

	loginForm := make(url.Values)
	loginPage.Find("input").Each(func(i int, s *goquery.Selection) {
		loginForm[s.AttrOr("name", "")] = []string{s.AttrOr("value", "")}
	})
	loginForm["username"] = []string{username}
	loginForm["password"] = []string{password}
	resp, err = c.hc.PostForm("https://swgoh.gg/accounts/login/", loginForm)
	if err != nil {
		return err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("swgoh.gg: unexpected status code %d: %v", resp.StatusCode, resp.Status)
	}
	// Logged in!
	c.authorized = true
	return nil
}
