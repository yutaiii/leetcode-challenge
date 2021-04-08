// hint: https://leetcode.com/problems/remove-element/discuss/599188/Easy-Go-Solution
func removeElement(nums []int, val int) int {
	count := 0
	for i, num := range nums {
		// valとして渡された値を上書きする
		nums[i-count] = nums[i]
		if num == val {
			count += 1
		}
	}
	return len(nums) - count
}