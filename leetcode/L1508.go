package main

import (
	"fmt"
	"math"
)

/*
1508. 子数组和排序后的区间和
给你一个数组 nums ，它包含 n 个正整数。你需要计算所有非空连续子数组的和，
并将它们按升序排序，得到一个新的包含 n * (n + 1) / 2 个数字的数组。
请你返回在新数组中下标为 left 到 right （下标从 1 开始）的所有数字和
（包括左右端点）。由于答案可能很大，请你将它对 10^9 + 7 取模后返回。

示例 1：
输入：nums = [1,2,3,4], n = 4, left = 1, right = 5
输出：13
解释：所有的子数组和为 1, 3, 6, 10, 2, 5, 9, 3, 7, 4 。将它们升序排序后，
我们得到新的数组 [1, 2, 3, 3, 4, 5, 6, 7, 9, 10] 。下标从 le = 1 到 ri = 5
的和为 1 + 2 + 3 + 3 + 4 = 13 。

示例 2：
输入：nums = [1,2,3,4], n = 4, left = 3, right = 4
输出：6
解释：给定数组与示例 1 一样，所以新数组为 [1, 2, 3, 3, 4, 5, 6, 7, 9, 10] 。
下标从 le = 3 到 ri = 4 的和为 3 + 3 = 6 。

示例 3：
输入：nums = [1,2,3,4], n = 4, left = 1, right = 10
输出：50
 */

func rangeSum(nums []int, n int, left int, right int) int {
	sumArray := make([]int, n + 1)
	MOD := int(math.Pow(10, 9) + 7)
	for i := 0; i < n; i++ {
		sumArray[i + 1] = sumArray[i] + nums[i]
	}
	// 计算前缀之和
	preSumArray := make([]uint64, n + 1)
	for i := 0; i < n; i++ {
		preSumArray[i + 1] = preSumArray[i] + uint64(sumArray[i + 1])
	}
	leftSum := presumForTarget(sumArray, preSumArray, left - 1, MOD)
	rightSum := presumForTarget(sumArray, preSumArray, right, MOD)
	return (rightSum - leftSum + MOD) % MOD
}

func presumForTarget(sumArray []int, preSumArray []uint64, target int, MOD int) int {
	sumTarget := binarySearcn1508(sumArray, target)
	output := 0
	for i, j := 0, 1; i < len(sumArray); i++ {
		for j < len(sumArray) && sumArray[j] - sumArray[i] < sumTarget {
			j++
		}
		/*
			sums[i]对应nums[0,i-1]
			如果[i,j-1)<sum,那么[i,i],...,[i,j-1)都<sum
			计算前缀和之和sum[i+1]-sum[i],....,sum[j-1]-sum[i]
			即sum[i+1]+...+sum[j-1]-sums[i]*(j-i-1)
			=preSums[j-1]-preSums[i]-sums[i]*(j-i-1)
		*/
		output = output + int(math.Mod(float64(preSumArray[j - 1] - preSumArray[i] - uint64(sumArray[i] * (j - i - 1))), float64(MOD)))
		output = output % MOD
		target = target - (j - i - 1)
	}
	return (output + target * sumTarget) % MOD
}

//二分连续子数组和，找到排在第target位的和
func binarySearcn1508(sumArray []int, target int) int {
	low, high := 0, sumArray[len(sumArray) - 1]
	for low < high {
		mid := (low + high) >> 1
		if countSum(sumArray, mid) < target {
			low = mid + 1
		} else if countSum(sumArray, mid) > target {
			high = mid
		} else {
			return mid
		}
	}
	return low
}

//连续子数组和<=sum的数量
func countSum(sumArray []int, target int) int {
	count := 0
	for i, j := 0, 1; i < len(sumArray); i++ {
		for j < len(sumArray) && sumArray[j] - sumArray[i] <= target {
			j++
		}
		count = count + (j - i - 1)
	}
	return count
}

func main() {
	fmt.Println(rangeSum([]int{1, 2, 3, 4}, 4, 1, 5))
}