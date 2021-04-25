package main

import (
	"container/heap"
	"fmt"
)

/*
1005. K 次取反后最大化的数组和
给定一个整数数组 A，我们只能用以下方法修改该数组：我们选择某个索引 i 并将 A[i] 替换为 -A[i]，然后总共重复这个过程 K 次。（我们可以多次选择同一个索引 i。）
以这种方式修改数组后，返回数组可能的最大和。
示例 1：
输入：A = [4,2,3], K = 1
输出：5
解释：选择索引 (1,) ，然后 A 变为 [4,-2,3]。

示例 2：
输入：A = [3,-1,0,2], K = 3
输出：6
解释：选择索引 (1, 2, 2) ，然后 A 变为 [3,1,0,2]

示例 3：
输入：A = [2,-3,-1,5,-4], K = 2
输出：13
解释：选择索引 (1, 4) ，然后 A 变为 [2,3,-1,5,4]。
 */

// 最小堆
type Heap1005 struct {
	heap []int
}

func (h Heap1005) Len() int {
	return len(h.heap)
}

func (h Heap1005) Less(i, j int) bool {
	return h.heap[i] < h.heap[j]
}

func (h Heap1005) Swap(i, j int) {
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}

func (h *Heap1005) Push(x interface{}) {
	h.heap = append(h.heap, x.(int))
}

func (h *Heap1005) Pop() interface{} {
	length := len(h.heap)
	number := h.heap[length - 1]
	h.heap = h.heap[:length - 1]
	return number
}

func (h *Heap1005) Sum() int {
	output := 0
	for i := 0; i < len(h.heap); i++ {
		output = output + h.heap[i]
	}
	return output
}

func largestSumAfterKNegations(A []int, K int) int {
	myHeap := &Heap1005{A}
	heap.Init(myHeap)
	for K > 0 {
		value := heap.Pop(myHeap).(int)
		if value < 0 {
			heap.Push(myHeap, -value)
			K--
		} else {
			if K % 2 == 1 {
				value = - value
			}
			heap.Push(myHeap, value)
			break
		}
	}
	return myHeap.Sum()
}

func main() {
	fmt.Println(largestSumAfterKNegations([]int{3,-1,0,2}, 3))
}

/*
解题思路
如果给定数组里有负数，那么反转负数是最优；反之选一个最小的正数不断反转即可，
 */