package main

import (
	"fmt"
	"sort"
)

/*
1424. 对角线遍历 II
给你一个列表 nums ，里面每一个元素都是一个整数列表。请你依照下面各图的规则，
按顺序返回 nums 中对角线上的整数。

示例 1：
输入：nums = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,4,2,7,5,3,8,6,9]

示例 2：
输入：nums = [[1,2,3,4,5],[6,7],[8],[9,10,11],[12,13,14,15,16]]
输出：[1,6,2,8,7,3,9,4,12,10,5,13,11,14,15,16]

示例 3：
输入：nums = [[1,2,3],[4],[5,6,7],[8],[9,10,11]]
输出：[1,4,2,5,3,8,6,9,7,10,11]

示例 4：
输入：nums = [[1,2,3,4,5,6]]
输出：[1,2,3,4,5,6]
 */

// 思路： 对下标排序
func findDiagonalOrder(nums [][]int) []int {
	indexArray := make([][2]int, 0)
	for i, row := range nums {
		for j, _ := range row {
			indexArray = append(indexArray, [2]int{i, j})
		}
	}
	sort.Slice(indexArray, func(i, j int) bool {
		a, b := indexArray[i], indexArray[j]
		sumI, sumJ := a[0] +a[1], b[0] +b[1]
		if sumI == sumJ {
			return a[0] > b[0]
		}
		return sumI < sumJ
	})
	output := make([]int, 0, len(indexArray))
	for _, index := range indexArray {
		output = append(output, nums[index[0]][index[1]])
	}
	return output
}

func main() {
	fmt.Println(findDiagonalOrder([][]int{{1,2,3,4,5}, {6, 7}, {8},
		{9, 10, 11}, {12, 13, 14, 15, 16}}))
}
