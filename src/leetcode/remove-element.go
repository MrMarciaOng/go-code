package leetcode



func removeElement(nums []int, val int) int {
 
	leftPointer := 0
	for i :=0; i < len(nums); i++{
			if nums[i] != val {
					nums[leftPointer] = nums[i]
					//inrecment left pointer
					leftPointer++
			}
	}
	return leftPointer
}