/*
Given an array of strings, group anagrams together.

Example:

Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
Output:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
Note:

All inputs will be in lowercase.
The order of your output does not matter.
*/

package group_anagrams

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	for _, str := range strs {
		b := []byte(str)
		sort.Slice(b, func(i, j int) bool {
			return b[i] > b[j]
		})
		m[string(b)] = append(m[string(b)], str)
	}

	res := make([][]string, 0, len(m))
	for _, s := range m {
		res = append(res, s)
	}

	return res
}
