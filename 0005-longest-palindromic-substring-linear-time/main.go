package main

import (
	"fmt"
	"strings"
)

// https://web.archive.org/web/20181208031518/https://articles.leetcode.com/longest-palindromic-substring-part-ii/#comment-7099
func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func preProcess(s string) string {
	if len(s) == 0 {
		return "^$"
	}

	var buf strings.Builder
	buf.WriteByte('^')
	for _, v := range s {
		buf.WriteByte('#')
		buf.WriteRune(v)
		//buf.WriteByte(byte(v))
	}
	buf.WriteString("#$")
	return buf.String()
}

func longestPalindrome(s string) string {
	var (
		t = preProcess(s)
		n = len(t) // length of t
		p = make([]int, n, n)
		c = 0 // center
		r = 0 // right edge
		maxLen      = 0
		centerIndex = 0
	)

	for i := 1; i < n-1; i++ {
		iMirror := 2*c - i // equals to i' = c - (i - c)
		if r > i {
			p[i] = min(r-i, p[iMirror])
		} else {
			p[i] = 0
		}

		// attempt to expand palindrome centered at i
		for t[i+1+p[i]] == t[i-1-p[i]] {
			p[i]++
		}

		// if palindrome centered at i expand past r
		// adjust center based on expanded palindrome
		if i+p[i] > r {
			c = i
			r = i + p[i]
		}

		if p[i] > maxLen {
			maxLen = p[i]
			centerIndex = i
		}
	}

	start := (centerIndex - 1 - maxLen) / 2
	return s[start : start+maxLen]
}

func main() {
	// abcbabcb
	// abcbabcb
	//s1 := "abcdbbfcba"
	s1 := "aaab"
	//  ^ # a # b # c # c # b # a # $
	// [0 0 1 0 1 0 1 6 1 0 1 0 1 0 0]
	fmt.Println(longestPalindrome(s1))
}
