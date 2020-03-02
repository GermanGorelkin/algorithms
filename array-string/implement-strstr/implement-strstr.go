/*
Implement strStr().

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:

Input: haystack = "hello", needle = "ll"
Output: 2
Example 2:

Input: haystack = "aaaaa", needle = "bba"
Output: -1
Clarification:

What should we return when needle is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().
*/

package implement_strstr

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	for i := range haystack {
		if len(haystack)-i >= len(needle) && haystack[i] == needle[0] {
			if needle == haystack[i:i+len(needle)] {
				return i
			}
		}
	}

	return -1
}

//func strStr(haystack string, needle string) int {
//	nLen := len(needle)
//	if nLen == 0 {
//		return 0
//	}
//	hLen := len(haystack)
//	if nLen > hLen {
//		return -1
//	}
//
//	var ok bool
//	for i := 0; i < hLen && (hLen-i)>=nLen; i++ {
//		if haystack[i] == needle[0] {
//			for j := 0; j < nLen; j++ {
//				ok = false
//				if (i+j) >= hLen || haystack[i+j] != needle[j] {
//					break
//				}
//				ok = true
//			}
//			if ok {
//				return i
//			}
//		}
//	}
//
//	return -1
//}
