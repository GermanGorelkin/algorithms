package open_the_lock

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_openLock(t *testing.T) {
	tests := map[string]struct {
		deadends []string
		target   string
		want     int
	}{
		"0": {deadends: []string{"0201", "0101", "0102", "1212", "2002"}, target: "0202", want: 6},
		"1": {deadends: []string{"8888"}, target: "0009", want: 1},
		"2": {deadends: []string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, target: "8888", want: -1},
		"3": {deadends: []string{"0000"}, target: "8888", want: -1},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, openLock(tc.deadends, tc.target))
		})
	}
}

func Test_incCh(t *testing.T) {
	tests := map[string]struct {
		input byte
		want1 byte
		want2 byte
	}{
		"0": {input: '0', want1: '9', want2: '1'},
		"1": {input: '1', want1: '0', want2: '2'},
		"2": {input: '2', want1: '1', want2: '3'},
		"9": {input: '9', want1: '8', want2: '0'},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			d1, d2 := incCh(tc.input)
			assert.Equal(t, tc.want1, d1)
			assert.Equal(t, tc.want2, d2)
		})
	}
}

func Test_genWheels(t *testing.T) {
	tests := map[string]struct {
		input string
		want  []string
	}{
		"0000": {input: "0000", want: []string{"9000", "1000", "0900", "0100", "0090", "0010", "0009", "0001"}},
		"9999": {input: "9999", want: []string{"8999", "0999", "9899", "9099", "9989", "9909", "9998", "9990"}},
		"1122": {input: "1122", want: []string{"0122", "2122", "1022", "1222", "1112", "1132", "1121", "1123"}},
	}

	wheels := make([]string, 8)

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			genWheels(tc.input, wheels)
			assert.Equal(t, tc.want, wheels)
		})
	}
}
