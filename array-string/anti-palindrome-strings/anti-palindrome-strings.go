package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var T int
	_, _ = fmt.Scanf("%d", &T)

	reader := bufio.NewReaderSize(os.Stdin, 200_000)

	for i := 0; i < T; i++ {
		line, _ := reader.ReadBytes('\n')
		l := len(line)
		if line[l-1] == '\n' {
			line = line[:len(line)-1]
			l--
		}

		isPalindrome := true
		for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
			if line[i] != line[j] {
				isPalindrome = false
				break
			}
		}
		if isPalindrome {
			fmt.Println(-1)
			continue
		}

		// count sort
		alph := make([]int, 26)
		for i := 0; i < l; i++ {
			alph[line[i]-'a']++
		}

		j := 0
		for i := 0; i < 26; i++ {
			for alph[i] > 0 {
				line[j] = byte(i) + 'a'
				alph[i]--
				j++
			}
		}

		fmt.Println(string(line))
	}
}
