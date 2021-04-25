package main

import (
	"fmt"
	"math"
)

/*
1491. 去掉最低工资和最高工资后的工资平均值
给你一个整数数组 salary ，数组里每个数都是 唯一 的，其中 salary[i] 是第 i 个
员工的工资。
请你返回去掉最低工资和最高工资以后，剩下员工工资的平均值。

示例 1：
输入：salary = [4000,3000,1000,2000]
输出：2500.00000
解释：最低工资和最高工资分别是 1000 和 4000 。
去掉最低工资和最高工资以后的平均工资是 (2000+3000)/2= 2500

示例 2：
输入：salary = [1000,2000,3000]
输出：2000.00000
解释：最低工资和最高工资分别是 1000 和 3000 。
去掉最低工资和最高工资以后的平均工资是 (2000)/1= 2000

示例 3：
输入：salary = [6000,5000,4000,3000,2000,1000]
输出：3500.00000

示例 4：
输入：salary = [8000,9000,2000,3000,6000,1000]
输出：4750.00000
 */

//func average(salary []int) float64 {
//	sort.Ints(salary)
//	length := len(salary)
//	if length <= 2 {
//		return 0
//	}
//	sum := 0
//	for i := 1; i < length - 1; i++ {
//		sum = sum + salary[i]
//	}
//	num := float64(sum / (length - 2))
//	value, _ := strconv.ParseFloat(fmt.Sprintf("%.5f", num), 64)
//	return value
//}

func average(salary []int) float64 {
	var output float64
	var max int
	min := math.MaxInt32
	for _, s := range salary {
		output = output + float64(s)
		if s > max {
			max = s
		}
		if s < min {
			min = s
		}
	}
	return (output - float64(max) - float64(min)) / float64(len(salary) - 2)
}

func main() {
	fmt.Println(average([]int{6000,5000,4000,3000,2000,1000}))
}