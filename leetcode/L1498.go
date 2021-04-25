package main

import (
	"fmt"
	"sort"
)

/*
1498. 满足条件的子序列数目
给你一个整数数组 nums 和一个整数 target 。
请你统计并返回 nums 中能满足其最小元素与最大元素的 和 小于或等于 target 的
非空 子序列的数目。
由于答案可能很大，请将结果对 10^9 + 7 取余后返回。
示例 1：
输入：nums = [3,5,6,7], target = 9
输出：4
解释：有 4 个子序列满足该条件。
[3] -> 最小元素 + 最大元素 <= target (3 + 3 <= 9)
[3,5] -> (3 + 5 <= 9)
[3,5,6] -> (3 + 6 <= 9)
[3,6] -> (3 + 6 <= 9)

示例 2：
输入：nums = [3,3,6,8], target = 10
输出：6
解释：有 6 个子序列满足该条件。（nums 中可以有重复数字）
[3] , [3] , [3,3], [3,6] , [3,6] , [3,3,6]

示例 3：
输入：nums = [2,3,3,4,6,7], target = 12
输出：61
解释：共有 63 个非空子序列，其中 2 个不满足条件（[6,7], [7]）
有效序列总数为（63 - 2 = 61）

示例 4：
输入：nums = [5,2,4,1,7,6,8], target = 16
输出：127
解释：所有非空子序列都满足条件 (2^7 - 1) = 127
 */

func numSubseq(nums []int, target int) int {
	MOD := 1000000007
	sort.Ints(nums)
	output := 0
	length := len(nums)
	if length == 0 {
		return 0
	}
	if nums[0] << 1 > target {
		return 0
	}
	powArray := make(map[int]int)
	powArray[0] = 1
	for i := 1; i < length; i++ {
		powArray[i] = (powArray[i - 1] * 2) % MOD
	}
	start, end := 0, length - 1
	for ; start < length; start++ {
		for ; end >= start; end-- {
			if nums[start] + nums[end] <= target {
				output = output + powArray[end - start]
				output = output % MOD
				break
			}
		}
		if end <= start {
			break
		}
	}
	return output
}

func main() {
	fmt.Println(numSubseq([]int{3, 3, 6, 8}, 10))
}

/*
//快速幂：防止long会溢出所以每一步都要mod
    private long myPow(int x, int n){
        if(n == 0) return 1;
        long t = x;
        long res = 1;
        while(n > 0){
            if((n & 1) == 1)
                res = (res * t) % (1000000000 + 7);
            t = (t * t) % (1000000000 + 7);
            n >>= 1;
        }
        return res % (1000000000 + 7);
    }
}
 */