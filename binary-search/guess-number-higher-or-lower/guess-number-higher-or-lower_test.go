package guess_number_higher_or_lower

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_guessNumber(t *testing.T) {
	tests := map[string]struct {
		pick int
		n    int
	}{
		"6 10":       {pick: 6, n: 10},
		"1 10":       {pick: 1, n: 10},
		"10 10":      {pick: 10, n: 10},
		"9999 10000": {pick: 9999, n: 10000},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			PICK = tc.pick
			assert.Equal(t, tc.pick, guessNumber(tc.n))
		})
	}
}
