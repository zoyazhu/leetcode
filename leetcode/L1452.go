package main

import (
	"fmt"
	"sort"
)

/*
1452. 收藏清单
给你一个数组 favoriteCompanies ，其中 favoriteCompanies[i] 是第 i 名用户收藏的
公司清单（下标从 0 开始）。
请找出不是其他任何人收藏的公司清单的子集的收藏清单，并返回该清单下标。下标需要按
升序排列。

示例 1：
输入：favoriteCompanies = [["leetcode","google","facebook"],
["google","microsoft"],["google","facebook"],["google"],["amazon"]]
输出：[0,1,4]
解释：
favoriteCompanies[2]=["google","facebook"] 是
favoriteCompanies[0]=["leetcode","google","facebook"] 的子集。
favoriteCompanies[3]=["google"] 是
favoriteCompanies[0]=["leetcode","google","facebook"] 和
favoriteCompanies[1]=["google","microsoft"] 的子集。
其余的收藏清单均不是其他任何人收藏的公司清单的子集，因此，答案为 [0,1,4] 。

示例 2：
输入：favoriteCompanies = [["leetcode","google","facebook"],
["leetcode","amazon"],["facebook","google"]]
输出：[0,1]
解释：favoriteCompanies[2]=["facebook","google"] 是
favoriteCompanies[0]=["leetcode","google","facebook"] 的子集，
因此，答案为 [0,1] 。

示例 3：
输入：favoriteCompanies = [["leetcode"],["google"],["facebook"],["amazon"]]
输出：[0,1,2,3]
 */

func peopleIndexes(favoriteCompanies [][]string) []int {
	for _, company := range favoriteCompanies {
		sort.Strings(company)
	}
	sortCompanies := make([][]string, len(favoriteCompanies))
	copy(sortCompanies, favoriteCompanies)
	sort.Slice(sortCompanies, func(i, j int) bool {
		return len(sortCompanies[i]) >= len(sortCompanies[j])
	})
	output := make([]int, 0, len(favoriteCompanies))
	for i, company := range favoriteCompanies {
		output = append(output, i)
		for _, sortCom := range sortCompanies {
			if len(sortCom) <= len(company) {
				break
			} else if isSubCompany(company, sortCom) {
				output = output[:len(output) - 1]
				break
			}
		}
	}
	return output
}

func isSubCompany(company, sortCompany []string) bool {
	lengthCom, lengthSort := len(company), len(sortCompany)
	com, sortCom := 0, 0
	for com < lengthCom && sortCom < lengthSort {
		if company[com] == sortCompany[sortCom] {
			com++
			sortCom++
		} else {
			sortCom++
		}
	}
	return com == lengthCom
}

func main() {
	fmt.Println(peopleIndexes([][]string{{"leetcode","google","facebook"},
		{"leetcode","amazon"}, {"facebook","google"}}))
}