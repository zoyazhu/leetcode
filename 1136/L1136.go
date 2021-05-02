package main

import "fmt"

/*
1136. 平行课程
已知有 N 门课程，它们以 1 到 N 进行编号。
给你一份课程关系表 relations[i] = [X, Y]，用以表示课程 X 和课程 Y 之间的先修关系：课程 X 必须在课程 Y 之前修完。
假设在一个学期里，你可以学习任何数量的课程，但前提是你已经学习了将要学习的这些课程的所有先修课程。
请你返回学完全部课程所需的最少学期数。
如果没有办法做到学完全部这些课程的话，就返回 -1。

示例 1：
输入：N = 3, relations = [[1,3],[2,3]]
输出：2
解释：
在第一个学期学习课程 1 和 2，在第二个学期学习课程 3。

示例 2：
输入：N = 3, relations = [[1,2],[2,3],[3,1]]
输出：-1
解释：
没有课程可以学习，因为它们相互依赖。
 */

func minimumSemesters(n int, relations [][]int) int {
	preClassNum := make(map[int]int)
	nextClasses := make([][]int, n)
	for i := 0; i < n; i++ {
		preClassNum[i] = 0
	}
	for _, class := range relations {
		preClassNum[class[1] - 1]++
		nextClasses[class[0] - 1] = append(nextClasses[class[0] - 1], class[1] - 1)
	}
	output := 0
	for len(preClassNum) > 0 {
		output++
		learned := make([]int, 0)
		for class, count := range preClassNum {
			if count == 0 {
				learned = append(learned, class)
			}
		}
		if len(learned) == 0 {
			return -1
		}
		for _, class := range learned {
			delete(preClassNum, class)
			for _, l := range nextClasses[class] {
				preClassNum[l]--
			}
		}
	}
	return output
}

func main() {
	fmt.Println(minimumSemesters(3, [][]int{{1, 3}, {2, 3}}))
}

/*
// 将 1 到 N，改成 0 到 N - 1，方便数组索引
func minimumSemesters(N int, relations [][]int) int {
    preClassCount := map[int]int{}
    nextClasses := make([][]int, N)
    for i := 0; i < N; i++ {
        preClassCount[i] = 0
    }
    for _, r := range relations {
        nextClasses[r[0]-1] = append(nextClasses[r[0]-1], r[1]-1) // 记录后置课程
        preClassCount[r[1]-1]++ // 计算初始入度
    }
    term := 0
    for len(preClassCount) > 0 {
        term++
        learn := []int{}
        for class, count := range preClassCount {
            if count == 0 {
                learn = append(learn, class) // 入度为 0，可以学习
            }
        }
        if len(learn) == 0 { // 没有课程可以学了，说明有循环
            return -1
        }
        for _, l := range learn {
            delete(preClassCount, l)
            for _, class := range nextClasses[l] {
                preClassCount[class]--
            }
        }
    }
    return term
}
 */