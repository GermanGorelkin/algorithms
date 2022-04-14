package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N int
	fmt.Scanf("%d ", &N)

	reader := bufio.NewReaderSize(os.Stdin, 10_000)
	matrix := ReadMatrix(reader, N)

	var even, odd int

	for i := 0; i < N; i++ {
		var edges int
		for j := 0; j < N; j++ {
			edges += matrix[i][j]
			if i == j {
				edges += matrix[i][j]
			}
		}
		if edges%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	fmt.Printf("%d %d\n", even, odd)
}

func ReadMatrix(reader *bufio.Reader, n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = ReadIntSlice(reader, n)
	}
	return matrix
}

func ReadIntSlice(reader *bufio.Reader, n int) []int {
	numbers := make([]int, n)

	var line []byte
	line, _ = reader.ReadSlice('\n')
	if line[len(line)-1] == '\n' {
		line = line[:len(line)-1]
	}
	if line[len(line)-1] == '\r' {
		line = line[:len(line)-1]
	}
	if line[len(line)-1] == ' ' {
		line = line[:len(line)-1]
	}
	ss := strings.Split(string(line), " ")

	for i, v := range ss {
		numbers[i], _ = strconv.Atoi(v)
	}

	return numbers
}
