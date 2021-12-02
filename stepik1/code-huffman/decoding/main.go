/*
https://stepik.org/lesson/13239/step/6?unit=3425
Задача на программирование: декодирование Хаффмана
*/

package main

import (
	"fmt"
)

func main() {
	var lenMap, lenCode int
	_, _ = fmt.Scanf("%d %d", &lenMap, &lenCode)
	codeMap := make(map[string]string, lenMap)

	for i := 0; i < lenMap; i++ {
		var ch, code string
		_, _ = fmt.Scanf("%s %s", &ch, &code)
		codeMap[code] = string(ch[0])
	}

	var code string
	_, _ = fmt.Scanf("%s", &code)

	res := decode(code, codeMap)
	fmt.Printf("%s", res)
}

func decode(code string, codeMap map[string]string) string {
	var res string
	i, j := 0, 1
	for j <= len(code) {
		if c, ok := codeMap[code[i:j]]; ok {
			res += c
			i, j = j, j+1
		} else {
			j++
		}
	}
	return res
}
