package main

import "fmt"

/*
1254. 统计封闭岛屿的数目
有一个二维矩阵 grid ，每个位置要么是陆地（记号为 0 ）要么是水域（记号为 1 ）。
我们从一块陆地出发，每次可以往上下左右 4 个方向相邻区域走，能走到的所有陆地区域，我们将其称为一座「岛屿」。
如果一座岛屿 完全 由水域包围，即陆地边缘上下左右所有相邻区域都是水域，那么我们将其称为 「封闭岛屿」。
请返回封闭岛屿的数目。

示例 1
输入：grid = [[1,1,1,1,1,1,1,0],[1,0,0,0,0,1,1,0],[1,0,1,0,1,1,1,0],[1,0,0,0,0,1,0,1],[1,1,1,1,1,1,1,0]]
输出：2
解释：
灰色区域的岛屿是封闭岛屿，因为这座岛屿完全被水域包围（即被 1 区域包围）。

示例 2：
输入：grid = [[0,0,1,0,0],[0,1,0,1,0],[0,1,1,1,0]]
输出：1

示例 3：
输入：grid = [[1,1,1,1,1,1,1],
             [1,0,0,0,0,0,1],
             [1,0,1,1,1,0,1],
             [1,0,1,0,1,0,1],
             [1,0,1,1,1,0,1],
             [1,0,0,0,0,0,1],
             [1,1,1,1,1,1,1]]
输出：2
 */

func closedIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	row := len(grid)
	column := len(grid[0])
	output := 0
	var dfs func(grid [][]int, i, j int) bool
	dfs = func(grid [][]int, i, j int) bool {
		if i < 0 || i >= row || j < 0 || j >= column {
			return false
		}
		if grid[i][j] == 1 {
			return true
		}
		grid[i][j] = 1
		left := dfs(grid, i - 1, j)
		right := dfs(grid, i + 1, j)
		top := dfs(grid, i, j - 1)
		bottom := dfs(grid, i, j + 1)
		return left && right && top && bottom
	}
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if grid[i][j] == 1 {
				continue
			}
			if dfs(grid, i, j) {
				output++
			}
		}
	}
	return output
}

func main() {
	fmt.Println(closedIsland([][]int{{1,1,1,1,1,1,1,0},{1,0,0,0,0,1,1,0},{1,0,1,0,1,1,1,0},{1,0,0,0,0,1,0,1},{1,1,1,1,1,1,1,0}}))
}

/*
func closedIsland(grid [][]int) int {
	num, rowNum, colNum := 0, len(grid), len(grid[0])
	for row, v1 := range grid {
		for col, v2 := range v1 {
			if v2 == 1 {
				continue
			}
			// 如果满足四面环水，则孤岛数量+1
			if dfs(grid, rowNum, colNum, row, col) {
				num++
			}
		}
	}

	return num
}

func dfs(grid [][]int, rowNum, colNum, row int, col int) (b bool) {
	// 如果触碰到边界了，无法构成孤岛
	if row < 0 || col < 0 || row >= rowNum || col >= colNum {
		return false
	}
	// 如果碰到水了，返回正确
	if grid[row][col] == 1 {
		return true
	}
	// 经过的地方置为海洋
	grid[row][col] = 1
	// 需要四个方向都环水
	up := dfs(grid, rowNum, colNum, row-1, col)
	down := dfs(grid, rowNum, colNum, row+1, col)
	left := dfs(grid, rowNum, colNum, row, col-1)
	right := dfs(grid, rowNum, colNum, row, col+1)
	return up && down && left && right
}
 */