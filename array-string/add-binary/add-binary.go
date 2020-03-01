/*
Given two binary strings, return their sum (also a binary string).

The input strings are both non-empty and contains only characters 1 or 0.

Example 1:

Input: a = "11", b = "1"
Output: "100"
Example 2:

Input: a = "1010", b = "1011"
Output: "10101
*/

package add_binary

/*
49 - 1
48 - 0
*/
const (
	zero byte = 48
	one  byte = 49
)

func addBinary(a string, b string) string {
	var n uint8
	var res []byte
	var r byte

	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		if i >= 0 {
			n += a[i]
		}
		if j >= 0 {
			n += b[j]
		}

		r, n = sum(n)
		res = append([]byte{r}, res...)
	}

	if n > 0 {
		res = append([]byte{one}, res...)
	}
	return string(res)
}

func sum(n uint8) (byte, uint8) {
	switch n {
	case zero:
		{
			return zero, 0
		}
	case one:
		{
			return one, 0
		}
	case zero + zero:
		{
			return zero, 0
		}
	case zero + one:
		{
			return one, 0
		}
	case one + one:
		{
			return zero, one
		}
	case one + zero + zero:
		{
			return one, 0
		}
	case one + one + one:
		{
			return one, one
		}
	case one + one + zero:
		{
			return zero, one
		}
	}
	return 0, 0
}
