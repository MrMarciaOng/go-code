package leetcode

import (
	"sort"
)

// IsAnagram checks if two strings are anagrams of each other
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sRuneArray := []rune(s)
	tRuneArray := []rune(t)
	// Sort both slices
	sort.Slice(sRuneArray, func(i, j int) bool { return sRuneArray[i] < sRuneArray[j] })
	sort.Slice(tRuneArray, func(i, j int) bool { return tRuneArray[i] < tRuneArray[j] })

	for i := 0; i < len(sRuneArray); i++ {
		if sRuneArray[i] != tRuneArray[i] {
			return false
		}
	}
	return true
}
