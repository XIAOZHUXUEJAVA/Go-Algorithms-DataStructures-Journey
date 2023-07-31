package string

func lengthOfLongestSubstring(s string) int {
	left := 0
	right := 0
	maxLength := 0
	winMap := map[byte]int{}
	for right < len(s) {
		a := s[right]
		right++
		winMap[a]++
		for winMap[a] > 1 {
			b := s[left]
			left++
			winMap[b]--
		}
		if maxLength < right-left {
			maxLength = right - left
		}
	}
	return maxLength
}
