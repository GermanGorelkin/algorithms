package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	h := heap{t: maxHeap}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		ss := strings.Split(s, " ")

		switch ss[0] {
		case "Insert":
			val, _ := strconv.Atoi(ss[1])
			h.Insert(val)
		case "ExtractMax":
			fmt.Printf("%d\n", h.Extract())
		}
	}
}

type heapType uint8

const (
	minHeap heapType = iota
	maxHeap
)

type heap struct {
	data []int
	t    heapType
}

func (h *heap) SiftDown(i int) {
	if h.t == minHeap {
		minSiftDown(h.data, i)
	} else {
		maxSiftDown(h.data, i)
	}
}

func minSiftDown(data []int, i int) {
	length := len(data)
	for 2*i+1 < length {
		left := 2*i + 1
		right := 2*i + 2

		smallest := left
		if right < length && data[right] < data[smallest] {
			smallest = right
		}

		if data[i] <= data[smallest] {
			break
		}

		data[smallest], data[i] = data[i], data[smallest]
		i = smallest
	}
}
func maxSiftDown(data []int, i int) {
	length := len(data)
	for 2*i+1 < length {
		left := 2*i + 1
		right := 2*i + 2

		largest := left
		if right < length && data[right] >= data[largest] {
			largest = right
		}

		if data[i] >= data[largest] {
			break
		}

		data[largest], data[i] = data[i], data[largest]
		i = largest
	}
}

func (h *heap) SiftUp(i int) {
	if h.t == minHeap {
		minSiftUp(h.data, i)
	} else {
		maxSiftUp(h.data, i)
	}
}
func minSiftUp(data []int, i int) {
	for data[i] < data[(i-1)/2] {
		data[i], data[(i-1)/2] = data[(i-1)/2], data[i]
		i = (i - 1) / 2
	}
}
func maxSiftUp(data []int, i int) {
	for data[i] > data[(i-1)/2] {
		data[i], data[(i-1)/2] = data[(i-1)/2], data[i]
		i = (i - 1) / 2
	}
}

func (h *heap) Insert(val int) {
	h.data = append(h.data, val)
	h.SiftUp(len(h.data) - 1)
}

func (h *heap) Extract() int {
	if len(h.data) == 0 {
		return 0
	}

	val := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.SiftDown(0)

	return val
}
