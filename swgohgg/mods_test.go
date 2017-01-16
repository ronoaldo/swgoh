package swgohgg

import (
	"reflect"
	"testing"
)

func TestModBonusSet(t *testing.T) {
	set := ModSet{
		"Transmitter": Mod{
			Level:    15,
			BonusSet: "Speed",
			PrimStat: ModStat{Stat: "Offense", Value: 5.88, IsPercent: true},
		},
		"Receiver": Mod{
			Level:    15,
			BonusSet: "Speed",
			PrimStat: ModStat{Stat: "Speed", Value: 30.0, IsPercent: false},
		},
		"Processor": Mod{
			Level:    15,
			BonusSet: "Speed",
			PrimStat: ModStat{Stat: "Defense", Value: 11.55, IsPercent: true},
		},
		"Holo-Array": Mod{
			Level:    15,
			BonusSet: "Speed",
			PrimStat: ModStat{Stat: "Critical Damage", Value: 36.0, IsPercent: false},
		},
		"Data-bus": Mod{
			Level:    15,
			BonusSet: "Health",
			PrimStat: ModStat{Stat: "Protection", Value: 24.0, IsPercent: true},
		},
		"Multiplexer": Mod{
			Level:    15,
			BonusSet: "Critical Damage",
			PrimStat: ModStat{Stat: "Potency", Value: 24.0, IsPercent: true},
		},
	}
	t.Log("Given a set with 4xSpeed, 1xHealth, 1xCritical Damage...")
	assert(t, "  It should return +10 speed for BonusForSet(Speed)", set.BonusForSet("Speed"), 10.0)
	assert(t, "  It should return +0 for BonusForSet(Critical Damage)", set.BonusForSet("Critical Damage"), 0.0)
	assert(t, "  It should return +0 for BonusForSet(Health)", set.BonusForSet("Health"), 0.0)

	set = ModSet{
		"Transmitter": Mod{
			Level:    15,
			BonusSet: "Health",
			PrimStat: ModStat{Stat: "Offense", Value: 5.88, IsPercent: true},
		},
		"Receiver": Mod{
			Level:    15,
			BonusSet: "Health",
			PrimStat: ModStat{Stat: "Speed", Value: 30.0, IsPercent: false},
		},
		"Processor": Mod{
			Level:    15,
			BonusSet: "Health",
			PrimStat: ModStat{Stat: "Defense", Value: 11.55, IsPercent: true},
		},
		"Holo-Array": Mod{
			Level:    14,
			BonusSet: "Health",
			PrimStat: ModStat{Stat: "Critical Damage", Value: 36.0, IsPercent: false},
		},
		"Data-bus": Mod{
			Level:    14,
			BonusSet: "Health",
			PrimStat: ModStat{Stat: "Protection", Value: 24.0, IsPercent: true},
		},
		"Multiplexer": Mod{
			Level:    15,
			BonusSet: "Critical Damage",
			PrimStat: ModStat{Stat: "Potency", Value: 24.0, IsPercent: true},
		},
	}
	t.Log("Given a set with 5xHealth (1 lvl 14), 1xCritical Damage...")
	assert(t, "  It should return +0 speed for BonusForSet(Speed)", set.BonusForSet("Speed"), 0.0)
	assert(t, "  It should return +0 for BonusForSet(Critical Damage)", set.BonusForSet("Critical Damage"), 0.0)
	assert(t, "  It should return +7.5 for BonusForSet(Health)", set.BonusForSet("Health"), 7.5)
}

func TestModStatIsBetter(t *testing.T) {
	testCases := []struct {
		Stat1 ModStat
		Stat2 ModStat
	}{
		{
			Stat1: ModStat{Stat: "Speed", Value: 30.0},
			Stat2: ModStat{Stat: "Speed", Value: 26.0},
		},
		{
			Stat1: ModStat{Stat: "Health", Value: 1.23, IsPercent: true},
			Stat2: ModStat{Stat: "Health", Value: 0.57, IsPercent: true},
		},
		{
			Stat1: ModStat{Stat: "Protection", Value: 300.0, IsPercent: false},
			Stat2: ModStat{Stat: "Protection", Value: 5.0, IsPercent: true},
		},
		{
			Stat1: ModStat{Stat: "Protection", Value: 1.0, IsPercent: false},
			Stat2: ModStat{Stat: "Protection", Value: 1.0, IsPercent: true},
		},
	}

	for _, tc := range testCases {
		t.Logf("Comparing stats: %v should be better than %v", tc.Stat1, tc.Stat2)
		if !tc.Stat1.IsBetterThan(tc.Stat2) {
			t.Errorf("Stat %v not better than %v, expected to be true", tc.Stat1, tc.Stat2)
		}
		if tc.Stat2.IsBetterThan(tc.Stat1) {
			t.Errorf("Stat %v better than %v, expected to be false", tc.Stat1, tc.Stat2)
		}
	}
}

func assert(t *testing.T, msg string, got, expected interface{}) {
	if reflect.DeepEqual(got, expected) {
		t.Log(msg + " PASS")
	} else {
		t.Errorf("%s: FAIL (Expected %#v[%T], got %#v[%T])", msg, expected, expected, got, got)
	}
}
