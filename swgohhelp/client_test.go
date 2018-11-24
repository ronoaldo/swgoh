package swgohhelp_test

import (
	"context"
	"flag"
	"testing"

	"github.com/ronoaldo/swgoh/swgohhelp"
)

var username, password, allyCode string

func init() {
	flag.StringVar(&username, "username", "ronoaldo", "Username to authenticate to the API")
	flag.StringVar(&password, "password", "", "Password to authenticate to the API")
	flag.StringVar(&allyCode, "ally", "335-983-287", "Ally code to run tests against")
}

func TestClientAuth(t *testing.T) {
	checkAuth(t, "ClientAuth")

	c := swgohhelp.New(context.Background()).SetDebug(true)
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
		player := players[i]
		t.Logf("Player %s (%d) %s", player.Name, player.AllyCode, player.Titles.Selected)
		for _, stat := range player.Stats {
			t.Logf("%s %d", stat.Name, stat.Value)
		}
		t.Logf("Arena rank %d", player.Arena.Char.Rank)
		t.Logf("Arena team: %v", player.Arena.Char.Squad)
		t.Logf("Ships rank %d", player.Arena.Ship.Rank)
		t.Logf("Ships team: %v", player.Arena.Ship.Squad)

		t.Logf("Roster: ")
		for _, unit := range player.Roster {
			t.Logf("%s %d* Lvl%d G%d", unit.Name, unit.Rarity, unit.Level, unit.Gear)
			if unit.Stats != nil {
				f := unit.Stats.Final
				m := unit.Stats.FromMods
				t.Logf("   %d (%d) speed, %d (%d) protection and %d (%d) health",
					f.Speed, m.Speed, f.Protection, m.Protection, f.Health, m.Health)
			}
		}
	}
}

func checkAuth(t *testing.T, name string) {
	if username == "" || password == "" {
		t.Skip("Missing credentials for test '%s'", name)
	}
}
