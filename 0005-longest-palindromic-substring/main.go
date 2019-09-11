package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func longestPalindrome(s string) string {
	if len(s) < 1 {
		return ""
	}

	var (
		start = 0
		end   = 0
		i     = 0
	)
	for ; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		length := max(len1, len2)
		if length > end-start {
			start = i - (length-1)/2
			end = i + length/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func main() {
	var r string
	s1 := "caba"
	r = longestPalindrome(s1)
	fmt.Printf("%q   --->    %q\n", s1, r)

	s2 := "babcbabcbaccba"
	r = longestPalindrome(s2)
	fmt.Printf("%q   --->    %q\n", s2, r)
}
