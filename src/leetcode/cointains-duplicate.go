package leetcode

func ContainsDuplicate(nums []int) bool {
	// Create a map to store numbers we've seen
	seenNumbers := make(map[int]int)

	// Iterate through each number in the array
	for _, currentNumber := range nums {
		// Check if number exists in map (more traditional syntax)
		numberExists := seenNumbers[currentNumber]
		if numberExists > 0 {
			return true // Found a duplicate
		}
		// If not seen, add it to our map
		seenNumbers[currentNumber] = 1
	}
	return false // No duplicates found
}
