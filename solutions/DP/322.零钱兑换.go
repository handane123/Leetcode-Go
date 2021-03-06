/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	// dp[i]: 凑足金额i,至少需要dp[i]种硬币
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1 // amount+1 means infinity
	}
	dp[0] = 0 // base case
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
// @lc code=end

