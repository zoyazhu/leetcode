package main

import "fmt"

/*
1273. 删除树节点
给你一棵以节点 0 为根节点的树，定义如下：
节点的总数为 nodes 个；
第 i 个节点的值为 value[i] ；
第 i 个节点的父节点是 parent[i] 。
请你删除节点值之和为 0 的每一棵子树。
在完成所有删除之后，返回树中剩余节点的数目。
示例 1：
输入：nodes = 7, parent = [-1,0,0,1,2,2,2], value = [1,-2,4,0,-2,-1,-1]
输出：2

示例 2：
输入：nodes = 7, parent = [-1,0,0,1,2,2,2], value = [1,-2,4,0,-2,-1,-2]
输出：6

示例 3：
输入：nodes = 5, parent = [-1,0,1,0,0], value = [-672,441,18,728,378]
输出：5

示例 4：
输入：nodes = 5, parent = [-1,0,0,1,1], value = [-686,-842,616,-739,-746]
输出：5
 */
func deleteTreeNodes(nodes int, parent []int, value []int) int {
	tree := make(map[int][]int)
	for i, v := range parent {
		node, ok := tree[i]
		if !ok {
			node = make([]int, 0)
		}
		if v != -1 {
			tree[i] = append(node, v)
		}
	}
	var dfs func(tree map[int][]int)
	output := 0
	dfs = func(tree map[int][]int) {

	}
	dfs(tree)
	return output
}

func main() {
	fmt.Println(deleteTreeNodes(7, []int{-1,0,0,1,2,2,2}, []int{1,-2,4,0,-2,-1,-1}))
}

/*
class Solution:
    def deleteTreeNodes(self, nodes: int, parent: List[int], value: List[int]) -> int:
        edges = {x: list() for x in range(nodes)}
        for x, p in enumerate(parent):
            if p != -1:
                edges[p].append(x)

        node_cnt = [1] * nodes

        def dfs(u):
            for v in edges[u]:
                dfs(v)
                value[u] += value[v]
                node_cnt[u] += node_cnt[v]
            if value[u] == 0:
                node_cnt[u] = 0

        dfs(0)
        return node_cnt[0]
 */