package sgame

import (
	"testing"
)

func Test_Ocean_Definition(t *testing.T) {
	tt := []struct {
		testN string

		x, y           int
		expP1x, expP1y int
		expP2x, expP2y int
	}{
		{"1", 0, 0, 0, 0, 10, 10},
		{"2", 5, 5, 5, 5, 15, 15},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.testN, func(t *testing.T) {
			// --- Given ---
			o := NewOcean(tc.x, tc.y)

			// --- When ---
			gotP1x, gotP1y, gotP2x, gotP2y := o.Definition()

			// --- Then ---
			if tc.expP1x != gotP1x {
				t.Errorf("expected P1x %d got %d", tc.expP1x, gotP1x)
			}
			if tc.expP1y != gotP1y {
				t.Errorf("expected P1y %d got %d", tc.expP1y, gotP1y)
			}
			if tc.expP2x != gotP2x {
				t.Errorf("expected P2x %d got %d", tc.expP2x, gotP2x)
			}
			if tc.expP2y != gotP2y {
				t.Errorf("expected P2y %d got %d", tc.expP2y, gotP2y)
			}
		})
	}
}
