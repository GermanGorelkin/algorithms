package utils

import (
	"bufio"
)

func ReadInt(r *bufio.Reader, buf []byte) (num int) {
	var err error
	var b byte
	buf = buf[:0]
	for {
		b, err = r.ReadByte()
		if b == ' ' || b == '\n' || err != nil {
			return parseInt(buf)
		} else {
			buf = append(buf, b)
		}
	}
}

func parseInt(b []byte) int {
	neg := false
	if b[0] == '-' {
		neg = true
		b = b[1:]
	}

	var n int
	for _, c := range b {
		c = c - '0'
		n *= 10
		n += int(c)
	}

	if neg {
		n = -n
	}
	return n
}
