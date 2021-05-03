package main

import "fmt"

/*
1245. 树的直径
给你这棵「无向树」，请你测算并返回它的「直径」：这棵树上最长简单路径的 边数。
我们用一个由所有「边」组成的数组 edges 来表示一棵无向树，其中 edges[i] = [u, v] 表示节点 u 和 v 之间的双向边。
树上的节点都已经用 {0, 1, ..., edges.length} 中的数做了标记，每个节点上的标记都是独一无二的。

示例 1：
输入：edges = [[0,1],[0,2]]
输出：2
解释：
这棵树上最长的路径是 1 - 0 - 2，边数为 2。

示例 2：
输入：edges = [[0,1],[1,2],[2,3],[1,4],[4,5]]
输出：4
解释：
这棵树上最长的路径是 3 - 2 - 1 - 4 - 5，边数为 4。
 */

func treeDiameter(edges [][]int) int {
	treeMap := make(map[int][]int)
	for _, edge := range edges {
		e1, e2 := edge[0], edge[1]
		if e1 > e2 {
			e1, e2 = e2, e1
		}
		node, ok := treeMap[e1]
		if !ok {
			node = make([]int, 0)
		}
		treeMap[e1] = append(node, e2)
	}
	var dfs func(treeMap map[int][]int, node int) (int, int)
	dfs = func(treeMap map[int][]int, node int) (int, int) {
		child, ok := treeMap[node]
		if !ok || child == nil {
			return 0, 0
		}
		diameter := 0
		leftHeight := -1
		rightHeight := -1
		for _, c := range child {
			h, d := dfs(treeMap, c)
			if diameter < d {
				diameter = d
			}
			if h > leftHeight {
				leftHeight, rightHeight = h, leftHeight
			} else if h > rightHeight {
				rightHeight = h
			}
		}
		if diameter < leftHeight + rightHeight + 2 {
			diameter = leftHeight + rightHeight + 2
		}
		return leftHeight + 1, diameter
	}
	_, output := dfs(treeMap, 0)
	return output
}

func main() {
	fmt.Println(treeDiameter([][]int{{0, 1}, {0, 2}}))
}

/*
思路：
因数组不好定位，故将 edges 转换为【有向树】tree，其中 v0 < v1，方向 v0->v1，Root = 0。
对 Root 递归做 DFS，计算树的直径
函数 DFS 定义: 返回当前树的高、当前树的直径
计算当前树高：取所有子树中最高的 h + 1
计算当前树直径，分 3 种情况讨论：
不存在子节点，返回 0
存在一个子节点，取当前树高。
存在两个子节点，设所有子树中最高为 h1，次高为 h2。则有 diameter = h1 + h2 + 2
PS: 4.2 和 4.3 两种情况，通过初始化 h1, h2 为 -1 简化代码实现。

func treeDiameter(edges [][]int) int {
    if edges == nil || len(edges) == 0 {
        return 0
    }
    tree := make(map[int][]int, len(edges))
    for _, v := range edges {
        v0, v1 := v[0], v[1]
        if v1 < v0 {
            v0, v1 = v[1], v[0]
        }
        nodes, ok := tree[v0]
        if !ok {
            nodes = []int{}
        }
        tree[v0] = append(nodes, v1)
    }
    _, path := dfs(tree, 0)
    return path
}

func dfs(tree map[int][]int, k int) (int, int) {
    childs, ok := tree[k]
    if !ok || len(childs) == 0 {
        return 0, 0
    }
    diameter, h1, h2 := 0, -1, -1
    for _, c := range childs {
        h, d := dfs(tree, c)
        if diameter < d {
            diameter = d
        }
        if h1 < h {
            h1, h2 = h, h1
        } else if h2 < h {
            h2 = h
        }
    }
    if diameter < h1 + h2 + 2 {
        diameter = h1 + h2 + 2
    }
    return h1 + 1, diameter
}
 */