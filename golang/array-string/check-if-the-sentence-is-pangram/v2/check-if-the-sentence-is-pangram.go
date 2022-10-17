/*
	https://leetcode.com/problems/check-if-the-sentence-is-pangram/

	A pangram is a sentence where every letter of the English alphabet appears at least once.
	Given a string sentence containing only lowercase English letters, return true if sentence is a pangram, or false otherwise.
*/

package check_if_the_sentence_is_pangram

/*
a - 1(1)
b - 2(10)
c - 4(100)
d - 8(1000)
....
z - 2^25
*/
func checkIfPangram(sentence string) bool {
	var seen uint64

	for _, ch := range sentence {
		seen |= (1 << (ch - 'a'))
	}

	return seen == (1<<26)-1
}
