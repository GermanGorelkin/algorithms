package v2

func firstUniqChar(s string) int {
	m := make([]int, 26)
	for i := 0; i < len(s); i++ {
		m[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if v := m[s[i]-'a']; v == 1 {
			return i
		}
	}
	return -1
}
