package main

import "fmt"

/*
797. 所有可能的路径
给一个有 n 个结点的有向无环图，找到所有从 0 到 n-1 的路径并输出（不要求
按顺序）
二维数组的第 i 个数组中的单元都表示有向图中 i 号结点所能到达的下一些结点（
译者注：有向图是有方向的，即规定了 a→b 你就不能从 b→a ）空就是没有下一
个结点了。

示例 1：
输入：graph = [[1,2],[3],[3],[]]
输出：[[0,1,3],[0,2,3]]
解释：有两条路径 0 -> 1 -> 3 和 0 -> 2 -> 3

示例 2：
输入：graph = [[4,3,1],[3,2,4],[3],[4],[]]
输出：[[0,4],[0,3,4],[0,1,3,4],[0,1,2,3,4],[0,1,4]]

示例 3：
输入：graph = [[1],[]]
输出：[[0,1]]

示例 4：
输入：graph = [[1,2,3],[2],[3],[]]
输出：[[0,1,2,3],[0,2,3],[0,3]]

示例 5：
输入：graph = [[1,3],[2],[3],[]]
输出：[[0,1,2,3],[0,3]]
 */
func allPathsSourceTarget(graph [][]int) [][]int {
	output := make([][]int, 0)
	var dfs func(current, end int, path []int)
	dfs = func(current, end int, path []int) {
		if current == end {
			copyPath := make([]int, len(path))
			copy(copyPath, path)
			output = append(output, append(copyPath, current))
		}
		for i := 0; i < len(graph[current]); i++ {
			dfs(graph[current][i], end, append(path, current))
		}
	}
	dfs(0, len(graph) - 1, []int{})
	return output
}

func main() {
	fmt.Println(allPathsSourceTarget([][]int{{1, 2}, {3}, {3}, {}}))
}

/*
实质求 A->B 两点之间的全部可达路径
每个一维数组为递归节点，数组长度为分支数，成员为下一个可达节点标识
建立相邻节点递归关系，遇到目标节点收录
 */