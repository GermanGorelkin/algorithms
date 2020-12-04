package contains_duplicate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_containsDuplicate(t *testing.T) {
	tests := map[string]struct {
		args []int
		want bool
	}{
		"[1,2,3,1]": {args: []int{1, 2, 3, 1}, want: true},
		"[1,2,3,4]": {args: []int{1, 2, 3, 4}, want: false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, containsDuplicate(tc.args))
		})
	}
}
