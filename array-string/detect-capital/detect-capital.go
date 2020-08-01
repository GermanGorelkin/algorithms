/*
Given a word, you need to judge whether the usage of capitals in it is right or not.

We define the usage of capitals in a word to be right when one of the following cases holds:

All letters in this word are capitals, like "USA".
All letters in this word are not capitals, like "leetcode".
Only the first letter in this word is capital, like "Google".
Otherwise, we define that this word doesn't use capitals in a right way.


Example 1:

Input: "USA"
Output: True


Example 2:

Input: "FlaG"
Output: False


Note: The input will be a non-empty word consisting of uppercase and lowercase latin letters.
*/

package detect_capital

func detectCapitalUse(word string) bool {
	if len(word) == 1 {
		return true
	}

	var min, max byte = 'a', 'z'
	if word[0] <= 'Z' {
		if word[1] <= 'Z' {
			min, max = 'A', 'Z'
		} else {
			min, max = 'a', 'z'
		}
	}

	for i := 1; i < len(word); i++ {
		if word[i] < min || word[i] > max {
			return false
		}
	}
	return true
}
