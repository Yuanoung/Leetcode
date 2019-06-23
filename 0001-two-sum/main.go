package main

import "fmt"

func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	for i, val := range nums {
		index, ok := dict[target-val]
		if ok {
			return []int{i, index}
		} else {
			dict[val] = i
		}
	}
	panic("don't exist")
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
