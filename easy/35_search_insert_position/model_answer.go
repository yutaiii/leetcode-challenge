// url: https://leetcode.com/problems/search-insert-position/
// hint:https://leetcode.com/problems/search-insert-position/discuss/732809/Go-recursive-binary-search-4ms
func searchInsert(nums []int, target int) int {
	// base case
	if len(nums) == 0 {
		return 0
	}

	mid := len(nums) / 2
	val := nums[mid]
	if val > target {
		// case 1: target in lower half
		return searchInsert(nums[:mid], target)
	}
	if val < target {
		// case 2: target in upper half
		// inc mid to exclude midpoint in re-slice
		mid++
		return mid + searchInsert(nums[mid:], target)
	}
	// case 3: val == target
	return mid
}