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
	edges := make([]Edge, 0, 1)

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if matrix[i][j] == 1 {
				edges = append(edges, Edge{From: i + 1, To: j + 1})
			}
		}
	}

	for _, e := range edges {
		fmt.Printf("%d %d\n", e.From, e.To)
	}
}

type Edge struct {
	From, To int
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
