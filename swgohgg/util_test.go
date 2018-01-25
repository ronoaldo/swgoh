package swgohgg

import "testing"

func TestSlug(t *testing.T) {
	testCases := []struct {
		in  string
		out string
	}{
		{"Ahsoka Tano (Fulcrum)", "ahsoka-tano-fulcrum"},
		{"Asajj Ventress", "asajj-ventress"},
		{"BB-8", "bb-8"},
		{"CC-2224 \"Cody\"", "cc-2224-cody"},
		{"Chirrut \u00cemwe", "chirrut-imwe"},
	}
	for i, tc := range testCases {
		t.Logf("Test case #%d", i)
		res := CharSlug(tc.in)
		t.Logf("Checking %s -> %s", tc.in, tc.out)
		if res != tc.out {
			t.Logf("FAIL: got %s", res)
		} else {
			t.Logf("OK")
		}
	}
}
