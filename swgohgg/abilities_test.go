package swgohgg

import "testing"

func TestAbilities(t *testing.T) {
	gg := NewClient("ronoaldo")
	zetas, err := gg.Zetas()
	if err != nil {
		t.Fatalf("Unexpected error when fetching zetas: %v", zetas)
	}

	for _, zeta := range zetas {
		t.Logf("Zeta returned: %s", zeta)
		if zeta.Name == "" || zeta.Character == "" {
			t.Errorf("Unexpected empty zeta metadata: %#v", zeta)
		}
	}
}
