/*
Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

Note: For the purpose of this problem, we define empty string as valid palindrome.

Example 1:

Input: "A man, a plan, a canal: Panama"
Output: true
Example 2:

Input: "race a car"
Output: false


Constraints:

s consists only of printable ASCII characters.
*/

package valid_palindrome

/*
65-90  A-Z
97-122 a-z
*/
func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	i, j := 0, len(s)-1
	for i < j {
		x := upper(s[i])
		if !isLetter(x) && !isNum(x) {
			i++
			continue
		}
		y := upper(s[j])
		if !isLetter(y) && !isNum(y) {
			j--
			continue
		}
		if x != y {
			return false
		}
		i++
		j--
	}

	return true
}

func isNum(x uint8) bool {
	if x >= '0' && x <= '9' {
		return true
	}
	return false
}

func isLetter(x uint8) bool {
	if (x >= 'A' && x <= 'Z') || (x >= 'a' && x <= 'z') {
		return true
	}
	return false
}

func upper(i uint8) uint8 {
	if i >= 97 && i <= 122 {
		return i - 32
	}
	return i
}
