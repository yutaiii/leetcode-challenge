// indexをメモしてあとでそれを取り除く方法
func removeElement(nums []int, val int) int {
	minusIndex := 0
	removeIndex := []int{}

	// 取り除くものをremoveIndexに追加
	for i, v := range nums {
		if v == val {
			removeIndex = append(removeIndex, i)
		}
	}

	// 配列から取り除く
	for i, _ := range nums {
		for _, v2 := range removeIndex {
			if i == v2 {
				// 要素を取り除くとindexの値が変わるので、その分indexをマイナスする
				nums = remove(nums, i+minusIndex)
				minusIndex--
			}
		}
	}
	return len(nums)
}

func remove(nums []int, i int) []int {
	// 取り除く要素が配列の最後の場合
	if len(nums) < i {
		return nums[:i-1]
	} else {
		// それ以外
		return append(nums[:i], nums[i+1:]...)
	}
}
