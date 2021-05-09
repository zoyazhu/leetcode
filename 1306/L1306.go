package main

import "fmt"

/*
1306. 跳跃游戏 III
这里有一个非负整数数组 arr，你最开始位于该数组的起始下标 start 处。当你位于下
标 i 处时，你可以跳到 i + arr[i] 或者 i - arr[i]。
请你判断自己是否能够跳到对应元素值为 0 的 任一 下标处。
注意，不管是什么情况下，你都无法跳到数组之外。

示例 1：
输入：arr = [4,2,3,0,3,1,2], start = 5
输出：true
解释：
到达值为 0 的下标 3 有以下可能方案：
下标 5 -> 下标 4 -> 下标 1 -> 下标 3
下标 5 -> 下标 6 -> 下标 4 -> 下标 1 -> 下标 3

示例 2：
输入：arr = [4,2,3,0,3,1,2], start = 0
输出：true
解释：
到达值为 0 的下标 3 有以下可能方案：
下标 0 -> 下标 4 -> 下标 1 -> 下标 3

示例 3：
输入：arr = [3,0,2,1,2], start = 2
输出：false
解释：无法到达值为 0 的下标 1 处。
*/

func canReach(arr []int, start int) bool {
	var dfs func(start int, arr []int, hasVisit []bool) bool
	hasVisit := make([]bool, len(arr))
	dfs = func(start int, arr []int, hasVisit []bool) bool {
		if start < 0 || start >= len(arr) || hasVisit[start] {
			return false
		}
		if arr[start] == 0 {
			return true
		}
		hasVisit[start] = true
		return dfs(start + arr[start], arr, hasVisit) ||
			dfs(start - arr[start], arr, hasVisit)
	}
	return dfs(start, arr, hasVisit)
}

func main() {
	fmt.Println(canReach([]int{4,2,3,0,3,1,2}, 5))
}