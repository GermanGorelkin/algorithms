/*
https://leetcode.com/problems/iterator-for-combination/
*/

package iterator_for_combination

type CombinationIterator struct {
	characters        string
	combinationLength int

	index int
	comb  []string
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	combIter := CombinationIterator{
		characters:        characters,
		combinationLength: combinationLength,
	}
	combIter.backtrack(characters, combinationLength, "", 0)

	return combIter
}

func (this *CombinationIterator) Next() string {
	s := this.comb[this.index]
	this.index++
	return s
}

func (this *CombinationIterator) HasNext() bool {
	return this.index < len(this.comb)
}

func (this *CombinationIterator) backtrack(s string, n int, combination string, i int) {
	if len(combination) == n {
		this.comb = append(this.comb, combination)
		return
	}

	for j := i; j < len(s); j++ {
		combination += string(s[j])
		this.backtrack(s, n, combination, j+1)
		combination = combination[:len(combination)-1]
	}
}
