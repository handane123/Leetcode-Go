/*
 * @lc app=leetcode.cn id=316 lang=golang
 *
 * [316] 去除重复字母
 */

// @lc code=start
func removeDuplicateLetters(s string) string {
	count := make([]int, 26) // 记录s中各字符出现次数
	seen := make([]bool, 26) // 当前字符是否出现在栈中
	var stack []byte // 单调递增栈，栈中存储的是字符

	// 统计各字符出现次数
	for i := range s {
		count[s[i]-'a']++
	}

	for i := range s {
		count[s[i]-'a']-- // 当前字符出现次数-1
		if !seen[s[i]-'a'] {
			// 当栈非空，栈顶元素比当前字符字典序大，且栈顶元素在后面还会出现，则出栈
			for len(stack) > 0 && stack[len(stack)-1] > s[i] && count[stack[len(stack)-1]-'a'] > 0 {
				seen[stack[len(stack)-1]-'a'] = false // 标记栈顶元素已不在栈中
				stack = stack[:len(stack)-1]			 
			}
			stack = append(stack, s[i]) // 当前字符入栈
			seen[s[i]-'a'] = true // 标记当前字符已入栈
		}	
	}		
	return string(stack) // 栈中元素即为结果
}
// @lc code=end

