package main

import "fmt"

func main() {
	ht := NewHashTable()
	var n int
	_, _ = fmt.Scanf("%d", &n)

	var cmd string
	var name string
	var number int
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%s", &cmd)

		switch cmd {
		case "add":
			{
				_, _ = fmt.Scanf("%d %s", &number, &name)
				ht.add(number, name)
			}
		case "del":
			{
				_, _ = fmt.Scanf("%d", &number)
				ht.del(number)
			}
		case "find":
			{
				_, _ = fmt.Scanf("%d", &number)
				name = ht.find(number)
				fmt.Println(name)
			}
		}
	}
}

const MAX_NUMBER = 10_000_000

func NewHashTable() *hashTable {
	return &hashTable{data: make([]string, MAX_NUMBER)}
}

type hashTable struct {
	data []string
}
func (ht *hashTable) add(number int, name string) {
	ht.data[number] = name
}
func (ht *hashTable) del(number int) {
	ht.data[number] = ""
}
func (ht *hashTable) find(number int) string {
	name := ht.data[number]
	if name == "" {
		return "not found"
	}
	return name
}