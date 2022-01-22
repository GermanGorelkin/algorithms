/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
Note:

All given inputs are in lowercase letters a-z.
*/

package longest_common_prefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	var pref []byte

	for i := 0; i < len(strs[0]); i++ {
		pref = append(pref, strs[0][i])
		for _, str := range strs[1:] {
			if len(str) <= i || strs[0][i] != str[i] {
				pref = pref[:len(pref)-1]
				return string(pref)
			}
		}
	}

	return string(pref)
}
