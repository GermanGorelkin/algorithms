package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N, M int
	fmt.Scanf("%d %d ", &N, &M)

	if N == 0 {
		fmt.Printf("%d\n", 0)
		return
	}

	list := ReadAdjacencyList(N, M)
	used := make([]int, N)

	var count int
	for {
		var exists bool
		for i, v := range used {
			if v == 0 {
				exists = true
				count++
				dfs(count, i, list, used)
			}
		}
		if !exists {
			break
		}
	}

	// fmt.Printf("%v\n\n", used)
	// fmt.Printf("%v\n\n", index)

	writer := bufio.NewWriterSize(os.Stdout, 100_000)

	writer.WriteString(strconv.Itoa(count))
	writer.WriteString("\n")

	comp := make([][]int, count)

	for i, v := range used {
		comp[v-1] = append(comp[v-1], i+1)
	}

	for _, vv := range comp {
		writer.WriteString(strconv.Itoa(len(vv)))
		writer.WriteString("\n")

		for _, v := range vv {
			writer.WriteString(strconv.Itoa(v))
			writer.WriteString(" ")
		}
		writer.WriteString("\n")
	}

	writer.Flush()
}

func dfs(con int, v int, list [][]int, used []int) {
	if used[v] != 0 {
		return
	}

	used[v] = con

	for _, v2 := range list[v] {
		dfs(con, v2, list, used)
	}
}

func ReadAdjacencyList(n, m int) [][]int {
	reader := bufio.NewReaderSize(os.Stdin, 1_000)
	list := make([][]int, n)

	for i := 0; i < m; i++ {
		from, to := ReadTwoInt(reader)
		from--
		to--

		list[from] = append(list[from], to)
		list[to] = append(list[to], from)
	}

	return list
}

func ReadTwoInt(reader *bufio.Reader) (from int, to int) {
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

	from, _ = strconv.Atoi(ss[0])
	to, _ = strconv.Atoi(ss[1])

	return
}
