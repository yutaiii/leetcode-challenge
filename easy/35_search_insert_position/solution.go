// url: https://leetcode.com/problems/search-insert-position/
func searchInsert(nums []int, target int) int {
	index := -1
	for i, v := range nums {
		if v == target {
			index = i
			break
		}

		// 中身がtargetより超えてしまう場合
		if v > target {
			index = i
			break
		}
	}

	// 配列の最後にtargetが入る場合
	if index == -1 {
		return len(nums)
	} else {
		return index
	}

}