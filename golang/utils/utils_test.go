package utils

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ReadInt(t *testing.T) {
	line := []byte("0 1 2 11 -123423123 2\n")
	r := bufio.NewReader(bytes.NewBuffer(line))
	var buf []byte
	assert.Equal(t, 0, ReadInt(r, buf))
	assert.Equal(t, 1, ReadInt(r, buf))
	assert.Equal(t, 2, ReadInt(r, buf))
	assert.Equal(t, 11, ReadInt(r, buf))
	assert.Equal(t, -123423123, ReadInt(r, buf))
	assert.Equal(t, 2, ReadInt(r, buf))
}

func Test_parseInt(t *testing.T) {
	tests := []struct {
		name string
		args []byte
		want int
	}{
		{name: "0", args: []byte("0"), want: 0},
		{name: "1", args: []byte("1"), want: 1},
		{name: "9", args: []byte("9"), want: 9},
		{name: "10", args: []byte("10"), want: 10},
		{name: "101", args: []byte("101"), want: 101},
		{name: "987654321", args: []byte("987654321"), want: 987654321},
		{name: "-987654321", args: []byte("-987654321"), want: -987654321},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, parseInt(tc.args))
		})
	}
}
