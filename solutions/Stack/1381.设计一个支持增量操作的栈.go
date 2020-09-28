/*
 * @lc app=leetcode.cn id=1381 lang=golang
 *
 * [1381] 设计一个支持增量操作的栈
 */

// @lc code=start
type CustomStack struct {
	stack []int
}


func Constructor(maxSize int) CustomStack {
	return CustomStack{stack: make([]int, 0, maxSize)}
}


func (this *CustomStack) Push(x int)  {
	if len(this.stack) == cap(this.stack) {
		return
	}
	this.stack = append(this.stack, x)
}


func (this *CustomStack) Pop() int {
	res := -1
	if len(this.stack) == 0 {
		return res
	}
	res = this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	return res
}


func (this *CustomStack) Increment(k int, val int)  {
	if k > len(this.stack) {
		k = len(this.stack)
	}
	for i := 0; i < k; i++ {
		this.stack[i] += val
	}
}


/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */
// @lc code=end

