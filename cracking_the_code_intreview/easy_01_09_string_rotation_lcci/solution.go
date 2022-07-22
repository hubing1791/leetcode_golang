package easy_01_09_string_rotation_lcci

import "strings"

//https://leetcode.cn/problems/string-rotation-lcci/
//2022-07-22
func isFlipedString(s1 string, s2 string) bool {
	n, m := len(s1), len(s2)
	if n != m {
		return false
	}
	if n == 0 {
		return true
	}
next:
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if s1[(i+j)%n] != s2[j] {
				continue next
			}
		}
		return true
	}
	return false
}

func isFlipedString2(s1 string, s2 string) bool {
	return len(s1) == len(s2) && strings.Contains(s1+s1, s2)
}
