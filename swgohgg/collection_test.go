package swgohgg

import (
	"testing"
)

func TestCharacterStats(t *testing.T) {
	c := NewClient("ronoaldo")
	stats, err := c.CharacterStats("tarkin")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Stats for 'tarkin': %#v", stats)
}
