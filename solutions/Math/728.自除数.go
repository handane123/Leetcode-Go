/*
 * @lc app=leetcode.cn id=728 lang=golang
 *
 * [728] 自除数
 */

// @lc code=start
func selfDividingNumbers(left int, right int) []int {
	var res []int
	for i := left; i <= right; i++ {
		if helper(i) {
			res = append(res, i)
		}
	}
	return res
}

func helper(num int) bool {
	if num < 10 {
		return true
	}
	origin := num
	for num > 0 {
		digit := num%10
		if digit == 0 || origin%digit != 0 {
			return false
		}
		num/=10
	}
	return true
}
// @lc code=end

