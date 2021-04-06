// index周りで動いていない
func removeElement(nums []int, val int) int {
	numsCopy := nums
	//minusIte := 0

	for i, v := range numsCopy {
		if v == val {
			// nums = remove(nums, i + minusIte)
			nums = remove(nums, i)
			//minusIte--
		}
	}
	return len(nums)
}

func remove(nums []int, i int) []int {
	if len(nums) < i {
		return nums[:i-1]
	} else {
		return append(nums[:i], nums[i+1:]...)
	}
}

// func remove(nums []int, i int) []int {
//     return append(nums[:i], nums[i+1:]...)
// }

// indexをメモしてあとでそれを取り除く方法