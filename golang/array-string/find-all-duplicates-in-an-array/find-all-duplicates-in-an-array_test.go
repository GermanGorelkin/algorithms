package find_all_duplicates_in_an_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findDuplicates(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  []int
	}{
		"1": {input: []int{4, 3, 2, 7, 8, 2, 3, 1}, want: []int{2, 3}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, findDuplicates(tc.input))
		})
	}
}
