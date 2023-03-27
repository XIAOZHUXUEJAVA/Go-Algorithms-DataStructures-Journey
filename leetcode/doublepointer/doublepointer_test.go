package doublepointer

import (
	"fmt"
	"testing"
)

func TestPalidromeNumber(t *testing.T) {

	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(1110))
	fmt.Println(isPalindrome(12321))

	fmt.Println(isPalindromeWithoutString(10))
	fmt.Println(isPalindromeWithoutString(1110))
	fmt.Println(isPalindromeWithoutString(12321))

}
