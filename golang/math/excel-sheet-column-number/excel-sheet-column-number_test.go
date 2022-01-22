package excel_sheet_column_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_titleToNumber(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		//"A":  {input: "A", want: 1},
		//"B":  {input: "B", want: 2},
		//"Z":  {input: "Z", want: 26},
		//"AA": {input: "AA", want: 27},
		//"AB": {input: "AB", want: 28},
		//"ZY": {input: "ZY", want: 701},
		"AAA": {input: "AAA", want: 703},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, titleToNumber(tc.input))
		})
	}
}
