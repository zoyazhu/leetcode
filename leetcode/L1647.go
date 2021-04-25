package main

import (
	"fmt"
	"sort"
)

/*
1647. 字符频次唯一的最小删除次数
如果字符串 s 中 不存在 两个不同字符 频次 相同的情况，就称 s 是 优质字符串 。
给你一个字符串 s，返回使 s 成为 优质字符串 需要删除的 最小 字符数。
字符串中字符的 频次 是该字符在字符串中的出现次数。例如，在字符串 "aab" 中，
'a' 的频次是 2，而 'b' 的频次是 1 。

示例 1：
输入：s = "aab"
输出：0
解释：s 已经是优质字符串。
示例 2：
输入：s = "aaabbbcc"
输出：2
解释：可以删除两个 'b' , 得到优质字符串 "aaabcc" 。
另一种方式是删除一个 'b' 和一个 'c' ，得到优质字符串 "aaabbc" 。

示例 3：
输入：s = "ceabaacb"
输出：2
解释：可以删除两个 'c' 得到优质字符串 "eabaab" 。
注意，只需要关注结果字符串中仍然存在的字符。（即，频次为 0 的字符会忽略不计。）
 */

func minDeletions(s string) int {
	countArray := make([]int, 26)
	for _, c := range s {
		countArray[c - 'a']++
	}
	output := 0
	sort.Ints(countArray)
	for i := len(countArray) - 2; i >= 0; i-- {
		if countArray[i] > 0 && countArray[i] >= countArray[i + 1] {
			target := max1647(0, countArray[i + 1] - 1)
			output = output + (countArray[i] - target)
			countArray[i] = target
		}
	}
	return output
}

func max1647(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minDeletions("aaabbbcc"))
}

/*
解题思路
1、 消除重复出现频率， 即需要先得到频率出现的频率。
2、 假如6次这个频率下有3个单词出现了。那么就需要看5,4,3,2,1这几个频率谁空闲。
然后分别把这两个单词删除至空闲频率。甚至完全删除。
3、 list01 := make([]int, liv[len(liv)-1]) 这么一个列表就记录了，那些频率空闲。
 */
