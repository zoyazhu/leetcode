package main

import (
	"fmt"
	"sort"
)

/*
1333. 餐厅过滤器
给你一个餐馆信息数组 restaurants，其中  restaurants[i] =
[idi, ratingi, veganFriendlyi, pricei, distancei]。你必须使用以下三个
过滤器来过滤这些餐馆信息。
其中素食者友好过滤器 veganFriendly 的值可以为 true 或者 false，如果为
true 就意味着你应该只包括 veganFriendlyi 为 true 的餐馆，为 false 则意
味着可以包括任何餐馆。此外，我们还有最大价格 maxPrice 和最大距离 maxDistance
两个过滤器，它们分别考虑餐厅的价格因素和距离因素的最大值。

过滤后返回餐馆的 id，按照 rating 从高到低排序。如果 rating 相同，那么按 id
从高到低排序。简单起见， veganFriendlyi 和 veganFriendly 为 true 时取值
为 1，为 false 时，取值为 0 。

示例 1：
输入：restaurants = [[1,4,1,40,10],[2,8,0,50,5],[3,8,1,30,4],[4,10,0,10,3],
[5,1,1,15,1]], veganFriendly = 1, maxPrice = 50, maxDistance = 10
输出：[3,1,5]
解释：
这些餐馆为：
餐馆 1 [id=1, rating=4, veganFriendly=1, price=40, distance=10]
餐馆 2 [id=2, rating=8, veganFriendly=0, price=50, distance=5]
餐馆 3 [id=3, rating=8, veganFriendly=1, price=30, distance=4]
餐馆 4 [id=4, rating=10, veganFriendly=0, price=10, distance=3]
餐馆 5 [id=5, rating=1, veganFriendly=1, price=15, distance=1]
在按照 veganFriendly = 1, maxPrice = 50 和 maxDistance = 10 进行过滤后，
我们得到了餐馆 3, 餐馆 1 和 餐馆 5（按评分从高到低排序）。

示例 2：
输入：restaurants = [[1,4,1,40,10],[2,8,0,50,5],[3,8,1,30,4],[4,10,0,10,3],[5,1,1,15,1]], veganFriendly = 0, maxPrice = 50, maxDistance = 10
输出：[4,3,2,1,5]
解释：餐馆与示例 1 相同，但在 veganFriendly = 0 的过滤条件下，应该考虑所有餐馆。
 */

func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	var output []int
	isVisit := 1
	var validArray [][]int
	for _, restaurant := range restaurants {
		isVisit = 1
		if veganFriendly == 1 {
			if restaurant[2] == 0 {
				isVisit = 0
			}
		}
		if restaurant[3] > maxPrice || restaurant[4] > maxDistance {
			isVisit = 0
		}
		if isVisit == 1 {
			validArray = append(validArray, restaurant)
		}
	}
	sort.Slice(validArray, func(i, j int) bool {
		if validArray[i][1] == validArray[j][1] {
			return validArray[i][0] > validArray[j][0]
		}
		return validArray[i][1] > validArray[j][1]
	})
	for _, valid := range validArray {
		output = append(output, valid[0])
	}
	return output
}

func main() {
	fmt.Println(filterRestaurants([][]int{{1,4,1,40,10}, {2,8,0,50,5},
		{3,8,1,30,4}, {4,10,0,10,3}, {5,1,1,15,1}}, 1, 50, 10))
}

/*
1.设置一个标志位 isadd 用来每次遍历元素 判断是否满足筛选条件 满足则设置为1，
否则设为0
 */