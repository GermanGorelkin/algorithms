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
	_, _ = fmt.Scanf("%d", &N)

	writer := bufio.NewWriterSize(os.Stdout, 1000_000)
	reader := bufio.NewReaderSize(os.Stdin, 1000_000)
	var line []byte
	var ss []string

	for i := 0; i < N; i++ {
		line, _ = reader.ReadSlice('\n')
		if line[len(line)-1] == '\n' {
			line = line[:len(line)-1]
		}
		ss = strings.Split(string(line), " ")
		N, _ := strconv.Atoi(ss[0])
		K, _ := strconv.Atoi(ss[1])

		line, _ = reader.ReadSlice('\n')
		if line[len(line)-1] == '\n' {
			line = line[:len(line)-1]
		}
		ss = strings.Split(string(line), " ")

		K = K % N

		for _, v := range ss[N-K:] {
			writer.WriteString(v)
			writer.WriteString(" ")
		}
		for _, v := range ss[:N-K] {
			writer.WriteString(v)
			writer.WriteString(" ")
		}
		writer.WriteString("\n")
	}
	writer.Flush()
}
