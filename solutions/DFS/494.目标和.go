/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */

// @lc code=start
func findTargetSumWays(nums []int, S int) int {
	// 深度优先搜索+备忘录
	type pair [2]int
	m := make(map[pair]int) // array in golang can be compared, so array can used as key in map
	var dfs func(i, remainder int) int
	dfs = func(i, remainder int) int {
		if i == len(nums) {
			if remainder == 0 {
				return 1
			}
			return 0
		}
		if val, ok := m[pair{i, remainder}]; ok {
			return val
		}
		ans := dfs(i+1, remainder+nums[i]) + dfs(i+1, remainder-nums[i])
		m[pair{i, remainder}] = ans
		return ans
	}
	return dfs(0, S)
}
// @lc code=end

