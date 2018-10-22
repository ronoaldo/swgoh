package swgohhelp_test

import (
	"context"
	"strings"
	"testing"

	"github.com/ronoaldo/swgoh/swgohhelp"
)

func TestDataPlayerTitles(t *testing.T) {
	checkAuth(t, "DataPlayerTitles")

	swapi := swgohhelp.New(context.Background()).SetDebug(true)
	if _, err := swapi.SignIn(username, password); err != nil {
		t.Fatalf("Unable to authorize client: %v", err)
	}

	titles, err := swapi.DataPlayerTitles()
	if err != nil {
		t.Fatalf("Unexpected error fetching titles: %v", titles)
	}

	for i := range titles {
		title := titles[i]
		t.Logf("Title #%s: %s -\n %s\n %s", i, title.Name, title.Desc, title.Details)
	}
}

func TestDataUnits(t *testing.T) {
	checkAuth(t, "TestDataUnits")

	swapi := swgohhelp.New(context.Background()).SetDebug(true)
	if _, err := swapi.SignIn(username, password); err != nil {
		t.Fatalf("Unable to authorize client: %v", err)
	}

	// Prepare caching
	t.Logf("Pre-fetching unit categories ...")
	swapi.DataUnitCategories()
	t.Logf("Pre-fetching unit skills ...")
	swapi.DataUnitSkills()
	t.Logf("Pre-feching unit abilities ...")
	swapi.DataUnitAbilities()

	units, err := swapi.DataUnits()
	if err != nil {
		t.Fatalf("Unable to fetch player units: %v", err)
	}

	for i := range units {
		unit := units[i]
		t.Logf("%s (%s)", unit.Name, strings.Join(unit.Categories, ","))
		t.Logf("> Type: %v (%v)", unit.CombatTypeName, unit.CombatType)
		t.Logf("> Skills")
		for _, skill := range unit.Skills {
			t.Logf("  %s (zeta: %v)", skill.Name, skill.IsZeta)
		}
		t.Logf("---")
	}
}
