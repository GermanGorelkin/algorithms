/*
	https://leetcode.com/problems/count-and-say/


*/

package count_and_say

import (
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	say := "1"

	for i := 1; i < n; i++ {
		say = count(say)
	}

	return say
}

func count(say string) string {
	var buf strings.Builder

	ch := say[0]
	n := 1
	for i := 1; i < len(say); i++ {
		if ch != say[i] {
			buf.WriteString(strconv.Itoa(n))
			buf.WriteByte(ch)

			n = 0
			ch = say[i]
		}
		n++
	}

	buf.WriteString(strconv.Itoa(n))
	buf.WriteByte(ch)

	return buf.String()
}
