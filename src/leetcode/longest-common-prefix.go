package leetcode

func LongestCommonPrefix(strs []string) string {
	// Handle edge cases
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	// Find shortest string first
	shortest := strs[0]
	for _, str := range strs {
		if len(str) < len(shortest) {
			shortest = str
		}
	}

	// Compare each character of shortest string with all other strings
	for i := range shortest {
		for _, str := range strs {
			if str[i] != shortest[i] {
				return shortest[:i]
			}
		}
	}

	return shortest
}
