package main

import (
	"container/list"
	"fmt"
)

func init() {
	calcDegreeX()
}

func main() {
	var n, m int
	_, _ = fmt.Scanf("%d", &m)
	_, _ = fmt.Scanf("%d", &n)

	ht := NewHashTable(m)
	var cmd string
	var word string
	var numBucket int
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%s", &cmd)

		switch cmd {
		case "add":
			{
				_, _ = fmt.Scanf("%s", &word)
				ht.add(word)
			}
		case "del":
			{
				_, _ = fmt.Scanf("%s", &word)
				ht.del(word)
			}
		case "find":
			{
				_, _ = fmt.Scanf("%s", &word)
				fmt.Println(ht.find(word))
			}
		case "check":
			{
				_, _ = fmt.Scanf("%d", &numBucket)
				bucket := ht.check(numBucket)
				for _, s := range bucket {
					fmt.Printf("%s ", s)
				}
				fmt.Printf("\n")
			}
		}
	}
}

type hashTable struct {
	data []*list.List
	m int
}
func NewHashTable(m int) *hashTable {
	ht := hashTable{
		data: make([]*list.List, m),
		m:    m,
	}
	for i := range ht.data {
		ht.data[i] = list.New()
	}
	return &ht
}
func (ht *hashTable) find(s string) string {
	h := getHash(s, ht.m)
	l := ht.data[h]

	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value.(string) == s {
			return FOUND_KEY
		}
	}
	return NOT_FOUND_KEY
}
func (ht *hashTable) del(s string) {
	h := getHash(s, ht.m)
	l := ht.data[h]

	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value.(string) == s {
			l.Remove(e)
		}
	}
}
func (ht *hashTable) add(s string) {
	h := getHash(s, ht.m)
	l := ht.data[h]

	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value.(string) == s {
			// s exists
			return
		}
	}

	l.PushFront(s)
}
func (ht *hashTable) check(i int) []string {
	var bucket []string
	l := ht.data[i]

	for e := l.Front(); e != nil; e = e.Next() {
		bucket = append(bucket, e.Value.(string))
	}

	return bucket
}

const (
	MAX_CHAR = 15
	X = 263
	P =  1_000_000_007

	FOUND_KEY = "yes"
	NOT_FOUND_KEY  = "no"
)

var (
	DegreeX = make([]uint64, MAX_CHAR)
)

func getHash(s string, m int) int {
	var hash uint64
	for i := 0; i < len(s); i++ {
		hash += uint64(s[i]) * DegreeX[i]
		hash %= P
	}
	hash %= uint64(m)

	if hash < 0 {
		hash += uint64(m)
	}

	return int(hash)
}

func calcDegreeX() {
	DegreeX[0] = 1
	DegreeX[1] = X

	for i := 2; i < MAX_CHAR; i++ {
		DegreeX[i] = DegreeX[i-1] * X % P
	}
}
