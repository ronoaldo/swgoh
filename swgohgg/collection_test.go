package swgohgg

import (
	"testing"
)

const (
	maxLevel     = 85
	maxGearLevel = 13
)

func TestCharacterStats(t *testing.T) {
	c := NewClient("ronoaldo")
	stats, err := c.CharacterStats("tarkin")
	if err != nil {
		t.Fatal(err)
	}

	if stats.Name != "Grand Moff Tarkin" {
		t.Errorf("Unexpected character name: '%s', expected 'Grand Moff Tarkin'", stats.Name)
	}
	if stats.Level < 1 || stats.Level > maxLevel {
		t.Errorf("Unexpected character level: %d, expected to be between [1, %d]", stats.Level, maxLevel)
	}
	if stats.GearLevel < 1 || stats.GearLevel > maxGearLevel {
		t.Errorf("Unexpected character gear: %d, expected to be between [1, %d]", stats.GearLevel, maxGearLevel)
	}
	if stats.Stars < 1 || stats.Stars > 7 {
		t.Errorf("Unexpected character stars: %d, expected to be between [1, 7]", stats.Stars)
	}
	if stats.GalacticPower < 1 {
		t.Errorf("Unexpected character power: %d, should be higher than zero", stats.GalacticPower)
	}

	if len(stats.Skills) != 6 {
		t.Errorf("Expected six character and pilot skills for Tarkin, got %v", stats.Skills)
	}

	if stats.Health < 1 {
		t.Errorf("Basic stats seems to be missing; got %d health", stats.Health)
	}

	t.Logf("Stats for 'tarkin': %#v", stats)
}
