/*
Given an arbitrary ransom note string and another string containing letters from all the magazines, write a function that will return true if the ransom note can be constructed from the magazines ; otherwise, it will return false.

Each letter in the magazine string can only be used once in your ransom note.

Note:
You may assume that both strings contain only lowercase letters.

canConstruct("a", "b") -> false
canConstruct("aa", "ab") -> false
canConstruct("aa", "aab") -> true
*/

package ransom_note

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) == 0 && len(magazine) == 0 {
		return true
	}
	if len(ransomNote) > len(magazine) {
		return false
	}

	hash := make([]int, 26)
	for i := 0; i < len(magazine); i++ {
		hash[magazine[i]-'a']++
	}
	for i := 0; i < len(ransomNote); i++ {
		hash[ransomNote[i]-'a']--
		if hash[ransomNote[i]-'a'] < 0 {
			return false
		}
	}
	return true
}

func canConstruct_v2(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	m := make(map[byte]int, len(magazine))
	for i := 0; i < len(magazine); i++ {
		m[magazine[i]]++
	}
	for i := 0; i < len(ransomNote); i++ {
		if v, ok := m[ransomNote[i]]; ok && v > 0 {
			m[ransomNote[i]]--
		} else {
			return false
		}
	}
	return true
}
