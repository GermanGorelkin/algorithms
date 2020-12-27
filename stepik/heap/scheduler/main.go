package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

type Item struct {
	priority int
	cpuNo    int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].priority == pq[j].priority {
		return pq[i].cpuNo < pq[j].cpuNo
	}
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func main() {
	var N, M int
	_, _ = fmt.Scanf("%d %d", &N, &M)

	pq := make(PriorityQueue, N)
	for i := 0; i < N; i++ {
		pq[i] = &Item{
			priority: 0,
			cpuNo:    i,
			index:    i,
		}
	}
	heap.Init(&pq)

	reader := bufio.NewReader(os.Stdin)
	for _, t := range ReadIntSlice(reader, M) {
		item := heap.Pop(&pq).(*Item)

		fmt.Printf("%d %d\n", item.cpuNo, item.priority)

		item.priority += t
		heap.Push(&pq, item)
	}
}

func ReadIntSlice(reader *bufio.Reader, n int) []int {
	numbers := make([]int, 0, n)

	var buf []byte
	var err error
	var b byte
	for {
		b, err = reader.ReadByte()
		if err != nil {
			break
		}
		if b == ' ' || b == '\n' {
			num, _ := strconv.Atoi(string(buf))
			numbers = append(numbers, num)

			if b == '\n' {
				break
			}
			buf = buf[:0]
		} else {
			buf = append(buf, b)
		}
	}

	return numbers
}
