package main

import (
	"fmt"
	"sort"
)

/*
1481. 不同整数的最少数目
给你一个整数数组 arr 和一个整数 k 。现需要从数组中恰好移除 k 个元素，
请找出移除后数组中不同整数的最少数目。

示例 1：
输入：arr = [5,5,4], k = 1
输出：1
解释：移除 1 个 4 ，数组中只剩下 5 一种整数。

示例 2：
输入：arr = [4,3,1,1,3,3,2], k = 3
输出：2
解释：先移除 4、2 ，然后再移除两个 1 中的任意 1 个或者三个 3 中的任意 1 个，
最后剩下 1 和 3 两种整数。
 */

type numStack struct {
	key int
	value int
}

func findLeastNumOfUniqueInts(arr []int, k int) int {
	numMap := make(map[int]int)
	for _, n := range arr {
		numMap[n]++
	}
	length := len(numMap)
	numStacks := make([]numStack, length)
	index := 0
	for key, value := range numMap {
		numStacks[index].key = key
		numStacks[index].value = value
		index++
	}
	sort.Slice(numStacks, func(i, j int) bool {
		return numStacks[i].value < numStacks[j].value
	})
	for _, n := range numStacks {
		if k >= n.value {
			k = k - n.value
			length--
		} else {
			break
		}
	}
	return length
}

func main() {
	fmt.Println(findLeastNumOfUniqueInts([]int{4,3,1,1,3,3,2}, 3))
}