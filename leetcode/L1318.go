package main

import "fmt"

/*
1318. 或运算的最小翻转次数
给你三个正整数 a、b 和 c。
你可以对 a 和 b 的二进制表示进行位翻转操作，返回能够使按位或运算   a OR b == c
成立的最小翻转次数。
「位翻转操作」是指将一个数的二进制表示任何单个位上的 1 变成 0 或者 0 变成 1 。

示例 1：
输入：a = 2, b = 6, c = 5
输出：3
解释：翻转后 a = 1 , b = 4 , c = 5 使得 a OR b == c

示例 2：
输入：a = 4, b = 2, c = 7
输出：1

示例 3：
输入：a = 1, b = 2, c = 3
输出：0
 */

func minFlips(a int, b int, c int) int {
	output := 0
	for {
		if a == 0 && b == 0 && c == 0 {
			break
		}
		if c % 2 == 0 {
			if a % 2 != 0 {
				output++
			}
			if b % 2 != 0 {
				output++
			}
		} else if a % 2 != 1 && b % 2 != 1 {
			output++
		}
		a = a / 2
		b = b / 2
		c = c / 2
	}
	return output
}

func main() {
	fmt.Println(minFlips(2, 6, 5))
}