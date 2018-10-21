package swgohhelp_test

import (
	"context"
	"encoding/json"
	"flag"
	"testing"

	"github.com/ronoaldo/swgoh/swgohhelp"
)

var username, password, allyCode string

func init() {
	flag.StringVar(&username, "username", "", "Username to authenticate to the API")
	flag.StringVar(&password, "password", "", "Password to authenticate to the API")
	flag.StringVar(&allyCode, "ally", "335-983-287", "Ally code to run tests against")
}

func TestClientAuth(t *testing.T) {
	checkAuth(t, "ClientAuth")

	c := swgohhelp.New(context.Background())
	token, err := c.SignIn(username, password)
	if err != nil {
		t.Fatalf("Unexpected error signing in: %v", err)
	}
	if token == "" {
		t.Fatalf("Unexpected empty token after auth!")
	}

	t.Log("Auth success!", token)
}

func TestPlayer(t *testing.T) {
	checkAuth(t, "Player")

	c := swgohhelp.New(context.Background()).SetDebug(true)
	if _, err := c.SignIn(username, password); err != nil {
		t.Fatalf("Unable to authorize client: %v", err)
	}
	players, err := c.Players(allyCode)
	if err != nil {
		t.Fatalf("Error fetching player: %v", err)
	}

	for i := range players {
		// Format for logging
		b, err := json.MarshalIndent(players[i], "", "  ")
		if err != nil {
			t.Errorf("> Unable to marshal player: %v", err)
			continue
		}
		t.Logf("players[%d] => %v (%s), updated at %v:\n%s", i, players[i].AllyCode, players[i].Name, players[i].UpdatedAt, string(b))
	}
}

func checkAuth(t *testing.T, name string) {
	if username == "" || password == "" {
		t.Fatalf("Missing credentials for test '%s'", name)
	}
}
