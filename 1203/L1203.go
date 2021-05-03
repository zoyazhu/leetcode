package main

import "fmt"

/*
1203. 项目管理
有 n 个项目，每个项目或者不属于任何小组，或者属于 m 个小组之一。group[i] 表示第 i 个项目所属的小组，如果第 i 个项目不属于任何小组，
则 group[i] 等于 -1。项目和小组都是从零开始编号的。可能存在小组不负责任何项目，即没有任何项目属于这个小组。
请你帮忙按要求安排这些项目的进度，并返回排序后的项目列表：

同一小组的项目，排序后在列表中彼此相邻。
项目之间存在一定的依赖关系，我们用一个列表 beforeItems 来表示，其中 beforeItems[i] 表示在进行第 i 个项目前（位于第 i 个项目左侧）
应该完成的所有项目。
如果存在多个解决方案，只需要返回其中任意一个即可。如果没有合适的解决方案，就请返回一个 空列表 。

示例 1：
输入：n = 8, m = 2, group = [-1,-1,1,0,0,1,0,-1], beforeItems = [[],[6],[5],[6],[3,6],[],[],[]]
输出：[6,3,4,1,5,2,0,7]

示例 2：
输入：n = 8, m = 2, group = [-1,-1,1,0,0,1,0,-1], beforeItems = [[],[6],[5],[6],[3],[],[4],[]]
输出：[]
解释：与示例 1 大致相同，但是在排序后的列表中，4 必须放在 6 的前面。
 */

func topoSort(graph [][]int, degree []int, items []int) []int {
	queue1 := make([]int, 0)
	output := make([]int, 0)
	for _, i := range items {
		if degree[i] == 0 {
			queue1 = append(queue1, i)
		}
	}
	for len(queue1) > 0 {
		first := queue1[0]
		queue1 = queue1[1:]
		output = append(output, first)
		for _, next := range graph[first] {
			degree[next]--
			if degree[next] == 0 {
				queue1 = append(queue1, next)
			}
		}
	}
	return output
}

func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
	groupItem := make([][]int, m+n)
	for i := range group {
		if group[i] == -1 {
			group[i] = m + i
		}
		groupItem[group[i]] = append(groupItem[group[i]], i)
	}
	graph := make([][]int, m + n)
	itemGraph := make([][]int, n)
	groupDegree := make([]int, m + n)
	itemDegree := make([]int, n)
	for i, item := range beforeItems {
		id := group[i]
		for _, pre := range item {
			preId := group[pre]
			// 不同的项目组，确定组间关系
			if preId != id {
				graph[preId] = append(graph[preId], id)
				groupDegree[id]++
			} else {
				itemGraph[pre] = append(itemGraph[pre], i)
				itemDegree[i]++
			}
		}
	}
	items := make([]int, m + n)
	for i := range items {
		items[i] = i
	}
	output := make([]int, 0)
	groupOrders := topoSort(graph, groupDegree, items)
	if len(groupOrders) < len(items) {
		return nil
	}
	for _, id := range groupOrders {
		item := groupItem[id]
		orders := topoSort(itemGraph, itemDegree, item)
		if len(orders) < len(item) {
			return nil
		}
		output = append(output, orders...)
	}
	return output
}

func main() {
	fmt.Println(sortItems(8, 2, []int{-1,-1,1,0,0,1,0,-1}, [][]int{{}, {6}, {5}, {6}, {3}, {}, {4}, {}}))
}

/*
func topSort(graph [][]int, deg, items []int) (orders []int) {
    q := []int{}
    for _, i := range items {
        if deg[i] == 0 {
            q = append(q, i)
        }
    }
    for len(q) > 0 {
        from := q[0]
        q = q[1:]
        orders = append(orders, from)
        for _, to := range graph[from] {
            deg[to]--
            if deg[to] == 0 {
                q = append(q, to)
            }
        }
    }
    return
}

func sortItems(n, m int, group []int, beforeItems [][]int) (ans []int) {
    groupItems := make([][]int, m+n) // groupItems[i] 表示第 i 个组负责的项目列表
    for i := range group {
        if group[i] == -1 {
            group[i] = m + i // 给不属于任何组的项目分配一个组
        }
        groupItems[group[i]] = append(groupItems[group[i]], i)
    }

    groupGraph := make([][]int, m+n) // 组间依赖关系
    groupDegree := make([]int, m+n)
    itemGraph := make([][]int, n) // 组内依赖关系
    itemDegree := make([]int, n)
    for cur, items := range beforeItems {
        curGroupID := group[cur]
        for _, pre := range items {
            preGroupID := group[pre]
            if preGroupID != curGroupID { // 不同组项目，确定组间依赖关系
                groupGraph[preGroupID] = append(groupGraph[preGroupID], curGroupID)
                groupDegree[curGroupID]++
            } else { // 同组项目，确定组内依赖关系
                itemGraph[pre] = append(itemGraph[pre], cur)
                itemDegree[cur]++
            }
        }
    }

    // 组间拓扑序
    items := make([]int, m+n)
    for i := range items {
        items[i] = i
    }
    groupOrders := topSort(groupGraph, groupDegree, items)
    if len(groupOrders) < len(items) {
        return nil
    }

    // 按照组间的拓扑序，依次求得各个组的组内拓扑序，构成答案
    for _, groupID := range groupOrders {
        items := groupItems[groupID]
        orders := topSort(itemGraph, itemDegree, items)
        if len(orders) < len(items) {
            return nil
        }
        ans = append(ans, orders...)
    }
    return
}
 */