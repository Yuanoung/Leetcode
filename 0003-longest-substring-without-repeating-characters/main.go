package main

import "fmt"

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

// sliding window
func lengthOfLongestSubstring(s string) int {
	n, ans := len(s), 0
	index := [128]int{}
	for j, i := 0, 0; j < n; j++ {
		i = max(index[s[j]], i)
		ans = max(ans, j-i+1)
		index[s[j]] = j + 1 // 这里为什么要用j+1？主要是为了i = max(...)。
	}
	return ans
}

// using hashmap
func usingHashMap(s string) int {
	n, ans := len(s), 0
	m := make(map[int]int)
	for i, j := 0, 0; j < n; j++ {
		if val, exist := m[int(s[j])]; exist {
			i = max(val, i)
		} else {
			ans = max(ans, j-i+1)
			m[int(s[j])] = j + 1
		}
	}
	return ans
}

func main() {
	s := "lllll"
	fmt.Println(lengthOfLongestSubstring(s))
	fmt.Println(usingHashMap(s))
}
