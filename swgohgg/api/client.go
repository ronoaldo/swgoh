package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

var errNotImplemented = fmt.Errorf("swgohggapi: not implemented")

// Client implements the API methods described by https://swgoh.gg/api/.
//
// This client wraps the default http.Client and exposes the API methods
// to make them easy to use, with some extra goodies.
// Data types returned can be easily manipulated and are designed to make
// coding with the API quick and easy tasks.
type Client struct {
	hc       http.Client
	ctx      context.Context
	endpoint string
	debug    bool
}

// NewClient initializes a new API client. This is required in order to prepare
// the internal client state, including the http.Client, context, logging and debuging.
func NewClient(c context.Context) *Client {
	return &Client{
		hc:       http.Client{},
		ctx:      c,
		endpoint: "https://swgoh.gg/api",
		debug:    false,
	}
}

// Debug enables debug info to stdout. This makes the cliente *very* verbose.
func (c *Client) Debug(debug bool) *Client {
	c.debug = debug
	return c
}

// get is an internal method make, log and debug HTTP calls to the API endpoint.
func (c *Client) get(path string, args ...interface{}) (resp *http.Response, err error) {
	url := fmt.Sprintf(c.endpoint+path, args...)
	resp, err = c.hc.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("swgohggapi: unexpected stauts code calling %s: %d %s", url, resp.StatusCode, resp.Status)
	}

	if c.debug {
		b, err := httputil.DumpResponse(resp, true)
		log.Printf("swgoggapi: GET %s => %v (err=%v)", path, string(b), err)
	}

	return resp, nil
}

// Characters returns a list of all game characters. Use this method
// to retrieve a list of characters and to use the resulting list in
// lookups and the like.
func (c *Client) Characters() (chars Characters, err error) {
	resp, err := c.get("/characters/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&chars); err != nil {
		return nil, err
	}

	return chars, nil
}

// Player returns the player detailed profile information including
// all player units.
func (c *Client) Player(allyCode string) (player *Player, err error) {
	resp, err := c.get("/player/%s", allyCode)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&player); err != nil {
		return nil, err
	}

	return player, nil
}

// Abilities list all available game abilities
func (c *Client) Abilities() (abilities []Ability, err error) {
	resp, err := c.get("/abilities/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&abilities); err != nil {
		return nil, err
	}

	return
}
