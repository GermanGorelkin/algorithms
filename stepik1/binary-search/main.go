package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line1, _ := reader.ReadString('\n')
	line2, _ := reader.ReadString('\n')
	line1 = strings.Trim(line1, "\n")
	line2 = strings.Trim(line2, "\n")

	var ss []string
	ss = strings.Split(line1, " ")
	data := make([]int, 0, len(ss)-1)
	for i := 1; i < len(ss); i++ {
		v, _ := strconv.Atoi(ss[i])
		data = append(data, v)
	}

	ss = strings.Split(line2, " ")
	for i := 1; i < len(ss); i++ {
		v, _ := strconv.Atoi(ss[i])
		indx := bsearch(data, v)
		fmt.Printf("%d ", indx)
	}
}

func bsearch(data []int, v int) int {
	lo, hi := 0, len(data)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if data[mid] == v {
			return mid + 1
		}
		if data[mid] > v {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return -1
}
