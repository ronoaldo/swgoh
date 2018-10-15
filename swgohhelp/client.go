package swgohhelp

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
)

var errNotImplemented = fmt.Errorf("swgohapi: not implemented")

// Client implements an authenticated callee to the https://api.swgoh.help service.
type Client struct {
	hc       *http.Client
	endpoint string
	token    string
	debug    bool
}

// New initializes an instance of Client making it ready to use.
func New(ctx context.Context) *Client {
	return &Client{
		hc:       http.DefaultClient,
		endpoint: "https://api.swgoh.help",
	}
}

// get is an internal method make, log and debug HTTP calls to the API endpoint.
func (c *Client) get(path string, args ...interface{}) (resp *http.Response, err error) {
	url := fmt.Sprintf(c.endpoint+path, args...)
	resp, err = c.hc.Get(url)
	if err != nil {
		return nil, err
	}

	if c.debug {
		b, err := httputil.DumpResponse(resp, true)
		log.Printf("swgohapi: GET %s => [err=%v] %v", path, string(b), err)
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("swgohapi: unexpected stauts code calling %s: %d %s", url, resp.StatusCode, resp.Status)
	}

	return resp, nil
}

func (c *Client) post(path, contentType string, body io.Reader, args ...interface{}) (resp *http.Response, err error) {
	url := fmt.Sprintf(c.endpoint+path, args...)
	resp, err = c.hc.Post(url, contentType, body)
	if err != nil {
		return nil, err
	}

	if c.debug {
		b, err := httputil.DumpResponse(resp, true)
		log.Printf("swgohapi: GET %s => [err=%v] %v", path, string(b), err)
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("swgohapi: unexpected stauts code calling %s: %d %s", url, resp.StatusCode, resp.Status)
	}

	return resp, nil
}

// SignIn authenticates the client and returns the accessToken or an error if authentication fails.
func (c *Client) SignIn(username, password string) (accessToken string, err error) {
	body := fmt.Sprintf("username=%s&password=%s&grant_type=password&client_id=goapiclient&client_secret=123456", username, password)
	resp, err := c.post("/auth/signin", "application/x-www-form-urlencoded", strings.NewReader(body))
	if err != nil {
		return "", err
	}

	var authResponse AuthResponse
	if err = json.NewDecoder(resp.Body).Decode(&authResponse); err != nil {
		return "", err
	}

	// Refresh token with the desired one
	c.token = authResponse.AccessToken

	return authResponse.AccessToken, nil
}

// Player retrieves the player profile stats
func (c *Client) Player(allyCodes ...string) (player []Player, err error) {
	return nil, errNotImplemented
}
