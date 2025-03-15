package leetcode

// [2,7,11,15]
// [3,2,4]
func twoSum(nums []int, target int) []int {
	// Create a map to store number->index pairs
	numMap := make(map[int]int)

	// One-pass solution using a hash map
	for i, num := range nums {
		complement := target - num
		if j, exists := numMap[complement]; exists {
			return []int{j, i}
		}
		numMap[num] = i
	}
	return []int{}
}
