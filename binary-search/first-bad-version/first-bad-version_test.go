package first_bad_version

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_firstBadVersion(t *testing.T) {
	tests := map[string]struct {
		pick int
		n    int
	}{
		"4 5":        {pick: 4, n: 5},
		"1 10":       {pick: 1, n: 10},
		"10 10":      {pick: 10, n: 10},
		"9999 10000": {pick: 9999, n: 10000},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			BAD_VESRION = tc.pick
			assert.Equal(t, tc.pick, firstBadVersion(tc.n))
		})
	}
}
