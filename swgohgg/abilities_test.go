package swgohgg

import "testing"

func TestAbilities(t *testing.T) {
	gg := NewClient("ronoaldo")
	zetas, err := gg.Zetas()
	t.Logf("zetas=%v; err=%v", zetas, err)
}
