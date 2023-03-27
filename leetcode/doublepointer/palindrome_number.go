package doublepointer

import "strconv"

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	for i, j := 0, len(str)-1; i < len(str)/2; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

func isPalindromeWithoutString(x int) bool {
	if x < 0 {
		return false
	}
	if x >= 0 && x < 10 {
		return true
	}
	if x%10 == 0 {
		return false
	}
	reverseNumber := 0

	for x > reverseNumber {
		reverseNumber = reverseNumber*10 + x%10
		x /= 10
	}
	return x == reverseNumber || x == reverseNumber/10

}
