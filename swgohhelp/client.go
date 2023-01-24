package swgohhelp

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"

	"github.com/ronoaldo/swgoh/cache"
)

//var errNotImplemented = fmt.Errorf("swgohhelp: not implemented")
// DefaultEndpoint is the default target host for API calls
var DefaultEndpoint = "https://api.swgoh.help"

var allyCodeCleanup = regexp.MustCompile("[^0-9]")

var targetEndpoint = flag.String("swgoh-endpoint", DefaultEndpoint, "Changes the default target endpoint to make calls to.")

var useExternalStats = flag.Bool("swgoh-external-statcalc", true, "Enables use of an external GP calculation API.")

// Client implements an authenticated callee to the https://api.swgoh.help service.
type Client struct {
	debug    bool
	endpoint string
	gameData cache.Cache
	guilds   cache.Cache
	hc       *http.Client
	players  cache.Cache
	token    string
}

// New initializes an instance of Client making it ready to use.
func New(ctx context.Context) *Client {
	client := &Client{
		hc:       http.DefaultClient,
		endpoint: *targetEndpoint,
	}
	cacheDir, err := cacheDirectory()
	if err != nil {
		log.Printf("swgohhelp: error loading cache directory: %v", err)
	}
	client.gameData = cache.NewCache(path.Join(cacheDir, GameDataCacheFile), GameDataCacheExpiration)
	client.players = cache.NewCache(path.Join(cacheDir, PlayerCacheFile), PlayerCacheExpiration)
	client.guilds = cache.NewCache(path.Join(cacheDir, GuildCacheFile), GuildCacheExpiration)
	return client
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
		return nil, fmt.Errorf("unable to create http request: %v", err)
	}

	if c.debug {
		b, _ := httputil.DumpRequestOut(req, true)
		writeLogFile(b, "req", method, urlPath)
	}

	resp, err = c.hc.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error when submitting http request: %v", err)
	}

	if c.debug {
		b, _ := httputil.DumpResponse(resp, true)
		writeLogFile(b, "resp", method, urlPath)
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("swgohhelp: unexpected status code calling %s: %d %s", url, resp.StatusCode, resp.Status)
	}

	return resp, nil
}

// SetDebug defines the debug state for the client.
func (c *Client) SetDebug(debug bool) *Client {
	c.debug = debug
	return c
}

// SignIn authenticates the client and returns the accessToken or an error if authentication fails.
func (c *Client) SignIn(username, password string) (accessToken string, err error) {
	body := fmt.Sprintf("username=%s&password=%s&grant_type=password&client_id=goapiclient&client_secret=123456", username, password)
	resp, err := c.call("POST", "/auth/signin", "application/x-www-form-urlencoded", strings.NewReader(body))
	if err != nil {
		return "", fmt.Errorf("issue when making sign in call: %v", err)
	}
	var authResponse AuthResponse
	if err = json.NewDecoder(resp.Body).Decode(&authResponse); err != nil {
		return "", fmt.Errorf("unable to decode auth response: %v", err)
	}
	// Refresh token with the desired one
	c.token = authResponse.AccessToken
	return authResponse.AccessToken, nil
}

// parseAllyCodes takes several ally code as strings and returns integer equivalents.
func parseAllyCodes(allyCodes ...string) (allyCodeNumbers []int, err error) {
	for _, a := range allyCodes {
		n, err := strconv.Atoi(allyCodeCleanup.ReplaceAllString(a, ""))
		if err != nil {
			return nil, err
		}
		allyCodeNumbers = append(allyCodeNumbers, n)
	}
	return
}

// writeLogFile is a debug helper function to write log data.
func writeLogFile(b []byte, reqresp, method, urlPath string) {
	urlPath = strings.Replace(urlPath, "/", "_", -1)
	fname := path.Join(os.TempDir(), fmt.Sprintf("swgohhelp%s-%s-%s.log", urlPath, method, reqresp))
	log.Printf("swgohhelp: writing log file %s: result: %v", fname, ioutil.WriteFile(fname, b, 0644))
}
