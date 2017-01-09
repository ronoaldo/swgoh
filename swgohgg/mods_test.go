package swgohgg

import (
	"testing"
)

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
