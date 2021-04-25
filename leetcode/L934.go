package main

import "fmt"

/*
934. 最短的桥
在给定的二维二进制数组 A 中，存在两座岛。（岛是由四面相连的 1 形成的一个最大组。）
现在，我们可以将 0 变为 1，以使两座岛连接起来，变成一座岛。
返回必须翻转的 0 的最小数目。（可以保证答案至少是 1 。）

示例 1：
输入：A = [[0,1],[1,0]]
输出：1

示例 2：
输入：A = [[0,1,0],[0,0,0],[0,0,1]]
输出：2

示例 3：
输入：A = [[1,1,1,1,1],[1,0,0,0,1],[1,0,1,0,1],[1,0,0,0,1],[1,1,1,1,1]]
输出：1
 */

type queue []*position
type position struct {
	i int
	j int
	depth int
}

func shortestBridge(A [][]int) int {
	size := len(A)
	if size < 1 {
		return 0
	}
	myQueue := newQueue()
	loop:
	for i := range A {
		for j := range A[i] {
			if A[i][j] == 1 {
				dfs934(A, i, j, myQueue)
				break loop
			}
		}
	}
	for {
		if len(*myQueue) < 1 {
			break
		}
		pos := myQueue.Dequeue()
		i, j, depth := pos.i, pos.j, pos.depth
		if i < 0 || j < 0 || i >= len(A) || j >= len(A[i]) ||
			(A[i][j] != 1 && A[i][j] != 0) {
			continue
		}
		if A[i][j] == 1 {
			return depth - 1
		}
		A[i][j] = -1
		myQueue.Enqueue(newPosition(i + 1, j, depth + 1))
		myQueue.Enqueue(newPosition(i - 1, j, depth + 1))
		myQueue.Enqueue(newPosition(i, j + 1, depth + 1))
		myQueue.Enqueue(newPosition(i, j - 1, depth + 1))
	}
	return -1
}

func dfs934(A [][]int, i, j int, queue *queue) {
	if i < 0 || j < 0 || i >= len(A) || j >= len(A[i]) {
		return
	}
	if A[i][j] == 0 {
		queue.Enqueue(newPosition(i, j, 1))
	} else if A[i][j] == 1 {
		A[i][j] = 2
		dfs934(A, i + 1, j, queue)
		dfs934(A, i - 1, j, queue)
		dfs934(A, i, j + 1, queue)
		dfs934(A, i, j - 1, queue)
	}
}

func newQueue() *queue {
	myQueue := make([]*position, 0)
	return (*queue)(&myQueue)
}

func newPosition(i, j, d int) *position {
	return &position{
		i: i,
		j: j,
		depth: d,
	}
}

func (q *queue) Enqueue(p *position) {
	*q = append(*q, p)
}

func (q *queue) Dequeue() *position {
	if len(*q) > 0 {
		p := (*q)[0]
		*q = (*q)[1:]
		return p
	}
	return nil
}

func main() {
	fmt.Println(shortestBridge([][]int{{0, 1}, {1, 0}}))
}