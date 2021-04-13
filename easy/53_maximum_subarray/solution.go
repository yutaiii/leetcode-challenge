// url: https://leetcode.com/problems/maximum-subarray/
// model answer: https://leetcode.com/problems/maximum-subarray/discuss/290007/Go-4-ms-100.00-easy-code
func maxSubArray(nums []int) int {
	// maxが起点のような使い方をされている
	max, sum := nums[0], nums[0]
	// nums[1:]は2番目から最後までを表す
	for _, v := range nums[1:] {
		// sumが0以下だったらsumを上書き
		if sum < 0 {
			sum = v
			// sumが0以上の時はsumに値を足す
		} else {
			sum += v
		}

		// maxがsumより小さい場合は値を書き換え
		if max < sum {
			max = sum
		}
	}
	return max
}