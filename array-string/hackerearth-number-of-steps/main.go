package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
// go build main.go && cat 430a731a427e11ea.txt.clean.txt | ./main
// Correct Output: 3056888
func main() {
	var N int
	_, _ = fmt.Scanf("%d", &N)

	reader := bufio.NewReaderSize(os.Stdin, 100_000)
	var line []byte
	var ss []string

	minVal := 5000 // max val
	// first line
	line, _ = reader.ReadSlice('\n')
	if line[len(line)-1] == '\n' {
		line = line[:len(line)-1]
	}
	ss = strings.Split(string(line), " ")

	a := make([]int, len(ss))
	for i, s := range ss {
		v, _ := strconv.Atoi(s)
		a[i] = v
		if v < minVal {
			minVal = v
		}
	}

	// second line
	line, _ = reader.ReadSlice('\n')
	if line[len(line)-1] == '\n' {
		line = line[:len(line)-1]
	}
	ss = strings.Split(string(line), " ")

	b := make([]int, len(ss))
	for i, s := range ss {
		v, _ := strconv.Atoi(s)
		b[i] = v
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