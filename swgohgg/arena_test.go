package swgohgg

import "testing"

func TestArena(t *testing.T) {
	gg := NewClient("ronoaldo")
	team, update, err := gg.Arena()
	if err != nil {
		t.Fatal(err)
	}
	for i := range team {
		char := team[i]
		t.Logf("Team member: %v", char)
	}
	t.Logf("Las update: %v", update)
}
