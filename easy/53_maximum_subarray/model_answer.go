// url: https://leetcode.com/problems/maximum-subarray/submissions/
// hint: https://leetcode.com/problems/maximum-subarray/discuss/309414/Go-Simple-DP-solution

func maxSubArray(nums []int) int {
	currSum, maxSum := nums[0], nums[0]
	for i := 0; i < len(nums); i++ {
		// 配列の数値か、今までの合計値のどちらが大きいかを比較する
		currSum = max(nums[i], currSum+nums[i])
		// 両者を比較　場合によってはmaxSumを置き換え
		maxSum = max(currSum, maxSum)
	}
	return maxSum
}

func max(a, b int) {
	if a > b {
		return a
	}
	return b
}