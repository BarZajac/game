package sgame

import (
	"testing"
)

func Test_nextOri(t *testing.T) {
	tt := []struct {
		testN string

		x, y       int
		w, h       int
		ori        Orientation
		dir        RotDir
		expX, expY int
		expOri     Orientation
	}{
		{"1", 0, 0, 4, 1, OriR, RotCW, 0, 0, OriD},
		{"2", 10, 10, 4, 1, OriR, RotCW, 10, 10, OriD},
		{"3", 11, 10, 4, 1, OriR, RotCCW, 11, 7, OriD},
		{"4", 11, 7, 1, 4, OriD, RotCW, 11, 10, OriR},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			gotX, gotY, gotOri := nextOri(tc.x, tc.y, tc.w, tc.h, tc.ori, tc.dir)

			// --- Then ---
			if gotX != tc.expX {
				t.Errorf("expected x %d got %d", tc.expX, gotX)
			}
			if gotY != tc.expY {
				t.Errorf("expected y %d got %d", tc.expY, gotY)
			}
			if gotOri != tc.expOri {
				t.Errorf("expected ori %s got %s", string(tc.expOri), string(gotOri))
			}
		})
	}
}
