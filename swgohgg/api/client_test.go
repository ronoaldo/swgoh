package api_test

import (
	"context"
	"flag"
	"strings"
	"testing"

	"github.com/ronoaldo/swgoh/swgohgg/api"
)

var allyCode string

func init() {
	flag.StringVar(&allyCode, "ally-code", "335983287", "The ally code to run tests against.")
}

func TestPlayerCall(t *testing.T) {
	c := api.NewClient(context.Background())
	player, err := c.Player(allyCode)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	t.Logf("Returned player profile: %#v, with %d units", player.PlayerProfile, len(player.Units))

	if player.Name != "Ronoaldo" {
		t.Errorf("Unexpected player name: %v for ally code %s", player.Name, allyCode)
	}

	if len(player.Units) == 0 {
		t.Errorf("Empty list of player units returned!")
	}

	// Assertions on Darth Vader
	dartVaderFound := false
	for _, unit := range player.Units {
		if unit.BaseID == "VADER" {
			dartVaderFound = true
			t.Logf("Vader found in unit list: %v", unit)

			if unit.Stats[api.Strength] == 0 {
				t.Logf("Vader Strength is zero!")
			}
		}
		if unit.Stats[api.PhysicalAccuracy] > 0 {
			t.Logf("> Unit with higher Accuracy: %v: %v", unit.BaseID, unit.Stats)
		}
	}
	if !dartVaderFound {
		t.Errorf("Darth Vader unit not found!")
	}
}

func TestCharactersCall(t *testing.T) {
	c := api.NewClient(context.Background())
	chars, err := c.Characters()
	if err != nil {
		t.Fatalf("Unexpected error listing characters: %v", err)
	}

	if len(chars) == 0 {
		t.Errorf("Empty character list returned")
	}

	// Assertions using Darth Vader
	darthVaderFound := false
	for _, char := range chars {
		if char.BaseID == "" {
			t.Errorf("Unexpected empty base_id field for character: %#v", char)
		}
		if char.Name == "Darth Vader" {
			darthVaderFound = true
			if char.PK == 0 {
				t.Errorf("Unexpected Darth Vader PK field: %d", char.PK)
			}
			if char.BaseID != "VADER" {
				t.Errorf("Unexpected Darth Vader base_id: %s, expected 'VADER'", char.BaseID)
			}
			if char.URL == "" || char.Image == "" {
				t.Errorf("Invalid endpoints for character and image: url=%v, image=%v", char.URL, char.Image)
			}
			if char.Power < 20000 {
				t.Errorf("Unexpected power for character: %d, expected a value higher than 20k", char.Power)
			}
			if char.Description == "" {
				t.Errorf("Empty character description: %s", char.Description)
			}
			if char.CombatType != 1 {
				t.Errorf("Unexpected combat type: %v, expected 1", char.CombatType)
			}
			if char.Alignment != "Dark Side" {
				t.Errorf("Come to the dark side, not '%s'", char.Alignment)
			}
			if len(char.AbilityClasses) == 0 {
				t.Errorf("Empty ability classes array")
			}
			for i := range char.AbilityClasses {
				if char.AbilityClasses[i] == "" {
					t.Errorf("Empty ability at position %d: %v", i, char.AbilityClasses[i])
				}
			}
			if char.Role != "Attacker" {
				t.Errorf("Unexpected role for Darth Vader: %s, expected 'Attacker'", char.Role)
			}
		}
	}
	if !darthVaderFound {
		t.Errorf("Missing 'Darth Vader' character!")
	}
}

func TestAbilitiesCall(t *testing.T) {
	c := api.NewClient(context.Background())
	abilities, err := c.Abilities()
	if err != nil {
		t.Fatalf("Unexpected error listing abilities: %v", err)
	}

	t.Logf("Decoded %d abilityes (err=%v)", len(abilities), err)

	if len(abilities) == 0 {
		t.Errorf("Invalid abilities length: 0")
	}

	for i := range abilities {
		ability := abilities[i]
		if strings.TrimSpace(ability.BaseID) == "" {
			t.Errorf("Empty ability base_id decoded: %v", ability)
		}
		if strings.TrimSpace(ability.Name) == "" {
			t.Errorf("Empty ability name decoded: %v", ability)
		}
	}
}
