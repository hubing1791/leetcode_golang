package medium_3_longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	hashtable := make(map[byte]int)
	n := len(s)
	pointer, ans := 0, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(hashtable, s[i-1])
		}
		for pointer < n && hashtable[s[pointer]] == 0 {
			hashtable[s[pointer]] += 1
			pointer += 1
		}
		if pointer-i > ans {
			ans = pointer - i
		}
	}
	return ans
}
