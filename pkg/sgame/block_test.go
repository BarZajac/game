package sgame

import (
	"testing"
)

func Test_nextOri(t *testing.T) {
	tt := []struct {
		testN string

		curr Orientation
		rot  RotDir
		exp  Orientation
	}{
		{"1", OriU, RotR, OriR},
		{"2", OriR, RotR, OriD},
		{"3", OriD, RotR, OriL},
		{"4", OriL, RotR, OriU},

		{"5", OriU, RotL, OriL},
		{"6", OriL, RotL, OriD},
		{"7", OriD, RotL, OriR},
		{"8", OriR, RotL, OriU},

		{"9", Orientation('Z'), RotR, Orientation('Z')},
		{"10", Orientation('Z'), RotL, Orientation('Z')},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			got := nextOri(tc.curr, tc.rot)

			// --- Then ---
			if got != tc.exp {
				t.Errorf("expected %s got %s", string(tc.exp), string(got))
			}
		})
	}
}
