package leetcode

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Keep track of position to place next unique element
	k := 1

	// Compare adjacent elements
	for i := 1; i < len(nums); i++ {
		//not the same as the previous one
		if nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}
