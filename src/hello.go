package main

import (
	"fmt"
	"go-code/leetcode"
)

func main() {
	// nums
	strs := []string{"dog", "racecar", "car"}
	// fmt.Println(leetcode.IsAnagram("anagram", "nagaram"))
	fmt.Println(leetcode.LongestCommonPrefix(strs))
}
