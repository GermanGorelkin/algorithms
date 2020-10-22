package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, m int

	_, _ = fmt.Scanf("%d", &n)

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.Trim(line, "\n")

	line2, _ := reader.ReadString('\n')
	line2 = strings.Trim(line2, "\n")
	ss := strings.Split(line2, " ")
	m, _ = strconv.Atoi(ss[0])

	q := maxQueue{}
	ss = strings.Split(line, " ")
	for i := 0; i < len(ss); i++ {
		v, _ := strconv.Atoi(ss[i])

		if i < m-1 {
			q.enQueue(v)
		} else {
			q.enQueue(v)
			fmt.Printf("%d ", q.max())
			q.deQueue()
		}
	}
}

type maxQueue struct {
	st1 maxStack
	st2 maxStack
}

func (q *maxQueue) enQueue(v int) {
	q.st1.push(v)
}
func (q *maxQueue) deQueue() int {
	if q.st2.isEmpty() {
		for !q.st1.isEmpty() {
			q.st2.push(q.st1.pop())
		}
	}
	return q.st2.pop()
}
func (q *maxQueue) max() int {
	return max(q.st1.max(), q.st2.max())
}


type value struct {
	val, max int
}
type maxStack struct {
	data []value
}
func (st *maxStack) push(v int) {
	val := value{val: v, max: v}
	if len(st.data) > 0 {
		val.max = max(st.data[len(st.data)-1].max, val.max)
	}
	st.data = append(st.data, val)
}
func (st *maxStack) pop() int {
	v := st.data[len(st.data)-1].val
	st.data = st.data[:len(st.data)-1]
	return v
}
func (st *maxStack) max() int {
	if len(st.data) > 0 {
		return st.data[len(st.data)-1].max
	}
	return -1
}
func (st *maxStack) isEmpty() bool {
	return len(st.data) == 0
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}