// url: https://leetcode.com/problems/maximum-subarray/
func maxSubArray(nums []int) int {
	minSub := 0
	sum := 0
	arr := []int{}

	for _, v := range nums {
		// マイナスの場合
		if v < 0 {
			// 初期値だったら値を更新
			if minSub == 0 {
				minSub = v
			}

			// 元のminSubと比較
			// 大きかったら書き換え
			if v > minSub {
				minSub = v
			}
		} else {
			// プラスの場合
			// 重複排除
			skipFlg := false
			for _, a := range arr {
				if v == a {
					skipFlg = true
				}
			}
			// 重複しているのものはskip
			if skipFlg {
				continue
			}
			// 重複していないものは合計値に追加
			sum += v
			arr = append(arr, v)
		}
	}
	return sum + minSub
}