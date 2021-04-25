package main

import "fmt"

/*
959. 由斜杠划分区域
在由 1 x 1 方格组成的 N x N 网格 grid 中，每个 1 x 1 方块由 /、\ 或空格
构成。这些字符会将方块划分为一些共边的区域。

（请注意，反斜杠字符是转义的，因此 \ 用 "\\" 表示。）。
返回区域的数目。

示例 1：
输入：
[
  " /",
  "/ "
]
输出：2
解释：2x2 网格如下：

示例 2：
输入：
[
  " /",
  "  "
]
输出：1
解释：2x2 网格如下：

示例 3：
输入：
[
  "\\/",
  "/\\"
]
输出：4
解释：（回想一下，因为 \ 字符是转义的，所以 "\\/" 表示 \/，而 "/\\" 表示 /\。）
2x2 网格如下：

示例 4：
输入：
[
  "/\\",
  "\\/"
]
输出：5
解释：（回想一下，因为 \ 字符是转义的，所以 "/\\" 表示 /\，而 "\\/" 表示 \/。）
2x2 网格如下：

示例 5：
输入：
[
  "//",
  "/ "
]
输出：3
解释：2x2 网格如下：
 */

type unionFind959 struct {
	parent []int
	size []int
	count int
}

func newUnionFind959(n int) *unionFind959 {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return &unionFind959{parent: parent, size: size, count: n}
}

func (u *unionFind959) find(x int) int {
	if u.parent[x] != x {
		u.parent[x] = u.find(u.parent[x])
	}
	return u.parent[x]
}

func (u *unionFind959) union(i, j int) {
	root1 := u.find(i)
	root2 := u.find(j)
	if root2 == root1 {
		return
	}
	if u.size[root1] < u.size[root2] {
		root1, root2 = root2, root1
	}
	u.size[root1] = u.size[root1] + u.size[root2]
	u.parent[root2] = root1
	u.count--
}

func regionsBySlashes(grid []string) int {
	length := len(grid)
	uf := newUnionFind959(4*length*length)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			 index := length * i + j
			 if i < length - 1 {
			 	bottom := index + length
			 	uf.union(index*4+2, bottom*4)
			 }
			 if j < length - 1 {
			 	right := index + 1
			 	uf.union(index*4+1, right*4+3)
			 }
			 if grid[i][j] == '/' {
			 	uf.union(index*4, index*4+3)
			 	uf.union(index*4+1, index*4+2)
			 } else if grid[i][j] == '\\' {
			 	uf.union(index*4, index*4+1)
			 	uf.union(index*4+2, index*4+3)
			 } else {
			 	uf.union(index*4, index*4+1)
			 	uf.union(index*4+1, index*4+2)
			 	uf.union(index*4+2, index*4+3)
			 }
		}
	}
	return uf.count
}

func main() {
	fmt.Println(regionsBySlashes([]string{"/\\", "\\/"}))
}
