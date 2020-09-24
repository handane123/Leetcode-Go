/*
 * @lc app=leetcode.cn id=1502 lang=golang
 *
 * [1502] 判断能否形成等差数列
 */

// @lc code=start
func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	for i := 1; i < len(arr)-1; i++ {
		if arr[i]*2 != arr[i-1]+arr[i+1] {
			return false
		}
	}
	return true
}
// @lc code=end

