/*
	https://leetcode.com/problems/check-if-the-sentence-is-pangram/

	A pangram is a sentence where every letter of the English alphabet appears at least once.
	Given a string sentence containing only lowercase English letters, return true if sentence is a pangram, or false otherwise.
*/

package check_if_the_sentence_is_pangram

func checkIfPangram(sentence string) bool {
	letters := [26]bool{}

	for _, ch := range sentence {
		letters[ch-'a'] = true
	}

	for _, ch := range letters {
		if !ch {
			return false
		}
	}

	return true
}
