package swgohhelp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"path"
	"strings"
)

var errNotImplemented = fmt.Errorf("swgohapi: not implemented")

// Client implements an authenticated callee to the https://api.swgoh.help service.
type Client struct {
	hc       *http.Client
	endpoint string
	token    string
	debug    bool
	cache    dataCache
}

// New initializes an instance of Client making it ready to use.
func New(ctx context.Context) *Client {
	return &Client{
		hc:       http.DefaultClient,
		endpoint: "https://api.swgoh.help",
	}
}

// SetDebug defines the debug state for the client.
func (c *Client) SetDebug(debug bool) *Client {
	c.debug = debug
	return c
}

// call internally makes and logs http requests to the API endpoints.
func (c *Client) call(method, urlPath, contentType string, body io.Reader, args ...interface{}) (resp *http.Response, err error) {
	url := fmt.Sprintf(c.endpoint+urlPath, args...)

	req, err := http.NewRequest(method, url, body)
	req.Header.Set("Content-type", contentType)
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}
	if err != nil {
		return nil, err
	}

	if c.debug {
		b, _ := httputil.DumpRequestOut(req, true)
		writeLogFile(b, "req", method, urlPath)
	}

	resp, err = c.hc.Do(req)
	if err != nil {
		return nil, err
	}

	if c.debug {
		b, _ := httputil.DumpResponse(resp, true)
		writeLogFile(b, "resp", method, urlPath)
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("swgohapi: unexpected stauts code calling %s: %d %s", url, resp.StatusCode, resp.Status)
	}

	return resp, nil
}

// SignIn authenticates the client and returns the accessToken or an error if authentication fails.
func (c *Client) SignIn(username, password string) (accessToken string, err error) {
	body := fmt.Sprintf("username=%s&password=%s&grant_type=password&client_id=goapiclient&client_secret=123456", username, password)
	resp, err := c.call("POST", "/auth/signin", "application/x-www-form-urlencoded", strings.NewReader(body))
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

// Players retrieves several player profile stats and roster details.
func (c *Client) Players(allyCodes ...string) (players []Player, err error) {
	payload, err := json.Marshal(map[string]interface{}{
		"allycode": allyCodes,
		"language": "eng_us",
		"enums":    true,
		"project": map[string]int{
			"id":         1,
			"allyCode":   1,
			"name":       1,
			"level":      1,
			"stats":      1,
			"arena":      1,
			"roster":     1,
			"guildName":  1,
			"guildRefId": 1,
			"titles":     1,
			"updated":    1,
		},
	})
	if err != nil {
		return nil, err
	}
	resp, err := c.call("POST", "/swgoh/player", "application/json", bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&players)
	if err != nil {
		return nil, err
	}

	// Enrich result with related data from data collections
	titles, err := c.DataPlayerTitles()
	if err != nil {
		return nil, err
	}
	for i := range players {
		players[i].Titles.Selected = titles[players[i].Titles.Selected].Name
		for j := range players[i].Titles.Unlocked {
			titleKey := players[i].Titles.Unlocked[j]
			players[i].Titles.Unlocked[j] = titles[titleKey].Name
		}
	}

	return players, nil
}

// DataPlayerTitles retrieves the data collection for player titles.
func (c *Client) DataPlayerTitles() (result map[string]DataPlayerTitle, err error) {
	if c.cache.playerTitles != nil {
		return c.cache.playerTitles, nil
	}
	// Prepare data collection call
	payload, err := json.Marshal(map[string]interface{}{
		"collection": "playerTitleList",
		"language":   "eng_us",
		"match": map[string]interface{}{
			"hidden":     false,
			"obtainable": true,
		},
	})
	if err != nil {
		return nil, err
	}
	// Parse result
	resp, err := c.call("POST", "/swgoh/data", "application/json", bytes.NewReader(payload))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	titles := make([]DataPlayerTitle, 0)
	err = json.NewDecoder(resp.Body).Decode(&titles)
	if err != nil {
		return
	}
	// Prepare response as map for easier usage
	result = make(map[string]DataPlayerTitle)
	for i := range titles {
		result[titles[i].ID] = titles[i]
	}
	c.cache.playerTitles = result
	return
}

// In memory cache of global game data.
type dataCache struct {
	playerTitles map[string]DataPlayerTitle
	units        map[string]Unit
}

// writeLogFile is a debug helper function to write log data.
func writeLogFile(b []byte, reqresp, method, urlPath string) {
	urlPath = strings.Replace(urlPath, "/", "_", -1)
	fname := path.Join(os.TempDir(), fmt.Sprintf("swgohhelp%s-%s-%s.log", urlPath, method, reqresp))
	log.Printf("swgohhelp: writing log file %s: result: %v", fname, ioutil.WriteFile(fname, b, 0644))
}
