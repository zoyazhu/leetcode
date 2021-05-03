package main

import "fmt"

/*
1192. 查找集群内的「关键连接」
力扣数据中心有 n 台服务器，分别按从 0 到 n-1 的方式进行了编号。
它们之间以「服务器到服务器」点对点的形式相互连接组成了一个内部集群，其中连接 connections 是无向的。
从形式上讲，connections[i] = [a, b] 表示服务器 a 和 b 之间形成连接。任何服务器都可以直接或者间接地通过网络到达任何其他服务器。
「关键连接」是在该集群中的重要连接，也就是说，假如我们将它移除，便会导致某些服务器无法访问其他服务器。
请你以任意顺序返回该集群内的所有 「关键连接」。

示例 1：
输入：n = 4, connections = [[0,1],[1,2],[2,0],[1,3]]
输出：[[1,3]]
解释：[[3,1]] 也是正确的。
 */

func criticalConnections(n int, connections [][]int) [][]int {
	criticalMap := make(map[int][]int)
	for _, connection := range connections {
		criticalMap[connection[0]] = append(criticalMap[connection[0]], connection[1])
		criticalMap[connection[1]] = append(criticalMap[connection[1]], connection[0])
	}
	output := make([][]int, 0)
	index := make([]int, n)
	for i := range index {
		index[i] = -1
	}
	var dfs func(node, currentIndex, prevNode int) int
	dfs = func(node, currentIndex, prevNode int) int {
		index[node] = currentIndex
		for _, next := range criticalMap[node] {
			if next == prevNode {
				continue
			}
			if index[next] == -1 {
				index[node] = min1192(index[node], dfs(next, currentIndex+1, node))
			} else {
				index[node] = min1192(index[node], index[next])
			}
		}
		if index[node] == currentIndex && index[node] != 0 {
			output = append(output, []int{prevNode, node})
		}
		return index[node]
	}
	dfs(0, 0, -1)
	return output
}

func min1192(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(criticalConnections(4, [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}}))
}
/*
解题思路
主要是通过标记节点的id值（发现顺序）找环 和 统计不在环上的边。
func criticalConnections(n int, connections [][]int) [][]int {
    // 建图
    m := make(map[int][]int)
    for _, c := range connections {
        m[c[0]] = append(m[c[0]], c[1])
        m[c[1]] = append(m[c[1]], c[0])
    }

    res := [][]int{}
    id := make([]int, n)
    for i := range id {
        id[i] = -1
    }
    // tarjan 思路 不断给节点标记id值(发现的先后顺序)， 遇到环就把换上节点统一到小值
    var dfs func(node, curId, perv int) int
    dfs = func(node, curId, perv int) int {
        id[node] = curId

        for _, next := range m[node] {
            if next == perv {
                continue
            }
            if id[next] == -1 {
                id[node] = min(id[node], dfs(next, curId+1, node))
            } else {
                id[node] = min(id[node], id[next])
            }
        }

        if id[node] == curId && node != 0 {
            res = append(res, []int{perv, node})
        }

        return id[node]
    }
    dfs(0, 0, -1)

    return res
}


func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
 */