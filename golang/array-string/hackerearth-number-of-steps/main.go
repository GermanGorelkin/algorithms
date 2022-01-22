package main

import (
	"bufio"
	"fmt"
	"os"
)
// go build main.go && cat 430a731a427e11ea.txt.clean.txt | ./main
// Correct Output: 3056888
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024)
	var buf []byte

	N := readInt(reader, buf)

	minVal := 5000 // max val
	// first line
	a := make([]int, N)
	for i := range a {
		a[i] = readInt(reader, buf)
		if a[i] < minVal {
			minVal = a[i]
		}
	}
	// second line
	b := make([]int, N)
	for i := range b {
		b[i] = readInt(reader, buf)
	}

	steps := numberOfSteps(a, b, minVal)

	fmt.Println(steps)
}

func numberOfSteps(a, b []int, min int) int {
	var steps int

	for i := 0; i < len(a); i++ {
		for a[i] > min && b[i] > 0 {
			a[i] -= b[i]
			steps++
		}
		if a[i] < min {
			min = a[i]
			i = 0
		}
		if a[i] < 0 || a[i] != min {
			steps = -1
			break
		}
	}

	return steps
}

// --
func readInt(r *bufio.Reader, buf []byte) (num int) {
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