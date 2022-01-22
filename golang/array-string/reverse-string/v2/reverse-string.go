package v2

func reverseString(s []byte) {
	reverse(s, 0, len(s)-1)
}

func reverse(s []byte, left, right int) {
	if left >= right {
		return
	}

	s[left], s[right] = s[right], s[left]
	right--
	left++

	reverse(s, left, right)
}
