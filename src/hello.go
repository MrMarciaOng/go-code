package main

import (
	"fmt"
	"go-code/leetcode"
)

func main() {
	// nums
	var nums = []int{1, 2, 3, 1}
	// fmt.Println(leetcode.IsAnagram("anagram", "nagaram"))
	fmt.Println(leetcode.ContainsDuplicate(nums))
}
