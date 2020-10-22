package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	_, _ = fmt.Scanf("%d", &n)

	st := maxStack{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		ss := strings.Split(s, " ")

		switch ss[0] {
		case "push":
			{
				val, _ := strconv.Atoi(ss[1])
				st.push(val)
			}
		case "pop":
			{
				st.pop()
			}
		case "max":
			{
				fmt.Printf("%d\n", st.max())
			}
		}
	}
}

type maxStack struct {
	data []int
	maxData []int
}
func (st *maxStack) push(v int) {
	st.data = append(st.data, v)

	if len(st.maxData) > 0 {
		v = max(st.maxData[len(st.maxData)-1], v)
	}
	st.maxData = append(st.maxData, v)
}
func (st *maxStack) pop() int {
	v := st.data[len(st.data)-1]
	st.data = st.data[:len(st.data)-1]
	st.maxData = st.maxData[:len(st.maxData)-1]
	return v
}
func (st *maxStack) max() int {
	return st.maxData[len(st.maxData)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}