package main

import "fmt"

/*
1640. 能否连接形成数组
给你一个整数数组 arr ，数组中的每个整数 互不相同 。另有一个由整数数组构成的数组
pieces，其中的整数也 互不相同 。请你以 任意顺序 连接 pieces 中的数组以
形成 arr 。但是，不允许 对每个数组 pieces[i] 中的整数重新排序。
如果可以连接 pieces 中的数组形成 arr ，返回 true ；否则，返回 false 。

示例 1：
输入：arr = [85], pieces = [[85]]
输出：true

示例 2：
输入：arr = [15,88], pieces = [[88],[15]]
输出：true
解释：依次连接 [15] 和 [88]


示例 3：
输入：arr = [49,18,16], pieces = [[16,18,49]]
输出：false
解释：即便数字相符，也不能重新排列 pieces[0]

示例 4：
输入：arr = [91,4,64,78], pieces = [[78],[4,64],[91]]
输出：true
解释：依次连接 [91]、[4,64] 和 [78]

示例 5：
输入：arr = [1,3,5,7], pieces = [[2,4,6,8]]
输出：false
 */

func canFormArray(arr []int, pieces [][]int) bool {
	formMap := make(map[int]int)
	for i, n := range arr {
		formMap[n] = i + 1
	}
	for i := 0; i < len(pieces); i++ {
		if len(pieces[i]) == 1 && formMap[pieces[i][0]] < 1 {
			return false
		}
		if len(pieces[i]) > 1 {
			for j := 0; j < len(pieces[i]) - 1; j++ {
				if formMap[pieces[i][j + 1]] - formMap[pieces[i][j]] != 1 {
					return false
				}
			}
		}
	}
	return true
}

func main() {
	fmt.Println(canFormArray([]int{15, 88}, [][]int{{88}, {15}}))
}