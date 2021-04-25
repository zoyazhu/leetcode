package main

import "fmt"

/*
1034. 边框着色
给出一个二维整数网格 grid，网格中的每个值表示该位置处的网格块的颜色。
只有当两个网格块的颜色相同，而且在四个方向中任意一个方向上相邻时，
它们属于同一连通分量。
连通分量的边界是指连通分量中的所有与不在分量中的正方形相邻（四个方向上）
的所有正方形，或者在网格的边界上（第一行/列或最后一行/列）的所有正方形。

给出位于 (r0, c0) 的网格块和颜色 color，使用指定颜色 color 为
所给网格块的连通分量的边界进行着色，并返回最终的网格 grid 。

示例 1：
输入：grid = [[1,1],[1,2]], r0 = 0, c0 = 0, color = 3
输出：[[3, 3], [3, 2]]

示例 2：
输入：grid = [[1,2,2],[2,3,2]], r0 = 0, c0 = 1, color = 3
输出：[[1, 3, 3], [2, 3, 3]]

示例 3：
输入：grid = [[1,1,1],[1,1,1],[1,1,1]], r0 = 1, c0 = 1, color = 2
输出：[[2, 2, 2], [2, 1, 2], [2, 2, 2]]
 */
func colorBorder(grid [][]int, r0 int, c0 int, color int) [][]int {
	if len(grid) == 0 {
		return grid
	}
	visited := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}
	currentColor := grid[r0][c0]
	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) {
			return 0
		}
		if visited[r][c] == true {
			return 1
		}
		if grid[r][c] != currentColor {
			return 0
		}
		visited[r][c] = true
		output := dfs(r + 1, c) + dfs(r - 1, c) + dfs(r, c - 1) + dfs(r, c + 1)
		if output < 4 {
			grid[r][c] = color
		}
		return 1
	}
	dfs(r0, c0)
	return grid
}

func main() {
	fmt.Println(colorBorder([][]int{{1, 1}, {1, 2}}, 0, 0, 3))
}