package main

import "fmt"

type packet struct {
	arrival, duration int
	start, finish     int
}

func main() {
	var size, n int

	_, _ = fmt.Scanf("%d %d", &size, &n)

	buf := NewRingBuffer(size)
	var arrival, duration int
	var prevComplete int

	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d %d", &arrival, &duration)
		p := packet{arrival: arrival, duration: duration}

		for !buf.isEmpty() {
			if pp := buf.front(); pp.finish <= p.arrival {
				buf.deQueue()
			} else {
				break
			}
		}

		if !buf.isFull() {
			if prevComplete < p.arrival {
				prevComplete = p.arrival
			}
			fmt.Printf("%d\n", prevComplete)

			prevComplete += p.duration
			p.finish = prevComplete
			buf.enQueue(p)

		} else {
			fmt.Printf("%d\n", -1)
		}
	}
}


func NewRingBuffer(len int) *queue {
	return &queue{
		data: make([]packet, len),
		len:  len,
		tail: -1,
		head: -1,
	}
}
type queue struct {
	data       []packet
	len        int
	head, tail int
}
func (q *queue) front() packet {
	return q.data[q.head]
}
func (q *queue) enQueue(v packet) {
	q.tail++
	q.tail %= q.len
	q.data[q.tail] = v

	if q.head == -1 {
		q.head = 0
	}
}
func (q *queue) deQueue() packet {
	v := q.data[q.head]

	if q.head == q.tail {
		q.head, q.tail = -1, -1
	} else {
		q.head++
		q.head %= q.len
	}

	return v
}
func (q *queue) isEmpty() bool {
	return q.head == -1
}
func (q *queue) isFull() bool {
	return ((q.tail + 1) % q.len) == q.head
}

/*
t
h
1 2 3

 */
