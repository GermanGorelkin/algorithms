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
	list := ReadReadAdjacencyListFromMatrix(reader, N)
	used := make([]int, N)

	var loopFound int

	for i := 0; i < N; i++ {
		if used[i] == 0 {
			if dfs(i, list, used) {
				loopFound = 1
				break
			}
		}
	}

	fmt.Printf("%d \n", loopFound)
}

/*
	0 - white
	1 - gray
	2 - black
*/

func dfs(v int, list [][]int, used []int) bool {
	used[v] = 1

	for _, v2 := range list[v] {
		if used[v2] == 0 {
			if dfs(v2, list, used) {
				return true
			}
		} else if used[v2] == 1 {
			return true
		}
	}

	used[v] = 2
	return false
}

func ReadReadAdjacencyListFromMatrix(reader *bufio.Reader, n int) [][]int {
	matrix := make([][]int, n)
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		ReadIntSlice(reader, slice)
		for j, v := range slice {
			if v == 1 {
				matrix[i] = append(matrix[i], j)
			}
		}
	}
	return matrix
}

func ReadIntSlice(reader *bufio.Reader, slice []int) {
	var line []byte
	line, _ = reader.ReadSlice('\n')
	if line[len(line)-1] == '\n' {
		line = line[:len(line)-1]
	}
	if line[len(line)-1] == ' ' {
		line = line[:len(line)-1]
	}
	ss := strings.Split(string(line), " ")

	for i, v := range ss {
		slice[i], _ = strconv.Atoi(v)
	}
}
