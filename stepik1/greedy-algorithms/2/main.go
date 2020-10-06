/*
Задача на программирование: непрерывный рюкзак

Первая строка содержит количество предметов 1<=n<=10^3 и вместимость рюкзака 0<=W<=2*10^6.
Каждая из следующих n строк задаёт стоимость 0<=c<=2*10^6
Выведите максимальную стоимость частей предметов (от каждого предмета можно отделить любую часть,
стоимость и объём при этом пропорционально уменьшатся),
помещающихся в данный рюкзак, с точностью не менее трёх знаков после запятой.

Sample Input:

3 50
60 20
100 50
120 30
Sample Output:

180.000
*/

package main

import (
	"fmt"
	"sort"
)

type item struct {
	c float64 // cost
	w float64 // weight
}

func main() {
	var maxWeight float64
	var count int
	_, _ = fmt.Scanf("%d %f", &count, &maxWeight)
	items := make([]item, count)

	for i := 0; i < count; i++ {
		var it item
		_, _ = fmt.Scanf("%f %f", &it.c, &it.w)
		items[i] = it
	}

	res := maxCost(items, maxWeight)

	fmt.Printf("%.3f", res)
}

func maxCost(items []item, weight float64) float64 {
	sort.SliceStable(items, func(i, j int) bool {
		return items[i].c/items[i].w > items[j].c/items[j].w
	})
	var maxCost float64

	for i := 0; i < len(items) && weight > 0; i++ {
		if items[i].w <= weight {
			maxCost += items[i].c
		} else {
			maxCost += items[i].c / items[i].w * weight
		}
		weight -= items[i].w
	}

	return maxCost
}
