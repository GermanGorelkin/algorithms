/*
Given a non-negative integer num represented as a string, remove k digits from the number so that the new number is the smallest possible.

Note:
The length of num is less than 10002 and will be ≥ k.
The given num does not contain any leading zero.
Example 1:

Input: num = "1432219", k = 3
Output: "1219"
Explanation: Remove the three digits 4, 3, and 2 to form the new number 1219 which is the smallest.
Example 2:

Input: num = "10200", k = 1
Output: "200"
Explanation: Remove the leading 1 and the number is 200. Note that the output must not contain leading zeroes.
Example 3:

Input: num = "10", k = 2
Output: "0"
Explanation: Remove all the digits from the number and it is left with nothing which is 0.
*/

package remove_k_digits

func removeKdigits(num string, k int) string {
	l := len(num)
	if l <= k {
		return "0"
	}
	st := &stack{}

	for i := 0; i < l; i++ {
		for !st.isEmpty() && st.peek() > num[i] && k > 0 {
			st.pop()
			k--
		}

		st.push(num[i])
	}

	for !st.isEmpty() && k > 0 {
		st.pop()
		k--
	}

	for len(st.data) > 1 && st.data[0] == '0' {
		st.data = st.data[1:]
	}

	return string(st.data)
}

type stack struct {
	data []byte
}

func (st *stack) push(v byte) {
	st.data = append(st.data, v)
}
func (st *stack) pop() byte {
	v := st.data[len(st.data)-1]
	st.data = st.data[:len(st.data)-1]
	return v
}
func (st *stack) isEmpty() bool {
	return len(st.data) == 0
}
func (st *stack) peek() byte {
	return st.data[len(st.data)-1]
}
