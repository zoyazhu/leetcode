package main

import "fmt"

/*
1310. 子数组异或查询
有一个正整数数组 arr，现给你一个对应的查询数组 queries，其中 queries[i] = [Li, Ri]。
对于每个查询 i，请你计算从 Li 到 Ri 的 XOR 值（即 arr[Li] xor arr[Li+1]
xor ... xor arr[Ri]）作为本次查询的结果。
并返回一个包含给定查询 queries 所有结果的数组。

示例 1：
输入：arr = [1,3,4,8], queries = [[0,1],[1,2],[0,3],[3,3]]
输出：[2,7,14,8]
解释：
数组中元素的二进制表示形式是：
1 = 0001
3 = 0011
4 = 0100
8 = 1000
查询的 XOR 值为：
[0,1] = 1 xor 3 = 2
[1,2] = 3 xor 4 = 7
[0,3] = 1 xor 3 xor 4 xor 8 = 14
[3,3] = 8
示例 2：
输入：arr = [4,8,2,10], queries = [[2,3],[1,3],[0,0],[0,3]]
输出：[8,0,4,4]
 */

func xorQueries(arr []int, queries [][]int) []int {
	preXor := make([]int, len(arr)+1)
	preXor[0] = 0
	for i := 0; i < len(arr); i++ {
		preXor[i + 1] = preXor[i] ^ arr[i]
	}
	output := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		output[i] = preXor[queries[i][0]] ^ preXor[queries[i][1]+1]
	}
	return output
}

func main() {
	fmt.Println(xorQueries([]int{1, 3, 4, 8}, [][]int{{0,1},{1,2},{0,3},{3,3}}))
}