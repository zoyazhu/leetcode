package main

import "fmt"

/*
1061. 按字典序排列最小的等效字符串
给出长度相同的两个字符串：A 和 B，其中 A[i] 和 B[i] 是一组等价字符。
举个例子，如果 A = "abc" 且 B = "cde"，那么就有 'a' == 'c', 'b' == 'd',
'c' == 'e'。
等价字符遵循任何等价关系的一般规则：
自反性：'a' == 'a'
对称性：'a' == 'b' 则必定有 'b' == 'a'
传递性：'a' == 'b' 且 'b' == 'c' 就表明 'a' == 'c'
例如，A 和 B 的等价信息和之前的例子一样，那么 S = "eed", "acd" 或 "aab"，
这三个字符串都是等价的，而 "aab" 是 S 的按字典序最小的等价字符串
利用 A 和 B 的等价信息，找出并返回 S 的按字典序排列最小的等价字符串。

示例 1：
输入：A = "parker", B = "morris", S = "parser"
输出："makkek"
解释：根据 A 和 B 中的等价信息，我们可以将这些字符分为 [m,p], [a,o],
[k,r,s], [e,i] 共 4 组。每组中的字符都是等价的，并按字典序排列。
所以答案是 "makkek"。

示例 2：
输入：A = "hello", B = "world", S = "hold"
输出："hdld"
解释：根据 A 和 B 中的等价信息，我们可以将这些字符分为 [h,w], [d,e,o],
[l,r] 共 3 组。所以只有 S 中的第二个字符 'o' 变成 'd'，最后答案为 "hdld"。

示例 3：
输入：A = "leetcode", B = "programs", S = "sourcecode"
输出："aauaaaaada"
解释：我们可以把 A 和 B 中的等价字符分为 [a,o,e,r,s,c], [l,p], [g,t]
和 [d,m] 共 4 组，因此 S 中除了 'u' 和 'd' 之外的所有字母都转化成了 'a'，
最后答案为 "aauaaaaada"。
 */

type unionSet struct {
	nums []int
	count int
}

func Constructor1061(n int) unionSet {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i
	}
	return unionSet{nums, n}
}

func (u *unionSet) find(n int) int {
	for u.nums[n] != n {
		n = u.nums[n]
	}
	return n
}

func (u *unionSet) union(i, j int) {
	rootI := u.find(i)
	rootJ := u.find(j)
	if rootI != rootJ {
		if rootI < rootJ {
			u.nums[rootJ] = rootI
		} else {
			u.nums[rootI] = rootJ
		}
	}
}

func smallestEquivalentString(A string, B string, S string) string {
	uf := Constructor1061(26)
	for i := 0; i < len(A); i++ {
		uf.union(int(A[i] - 'a'), int(B[i] - 'a'))
	}
	output := make([]byte, len(S))
	for i := 0; i < len(S); i++ {
		c := uf.find(int(S[i] - 'a'))
		output[i] = byte('a' + c)
	}
	return string(output)
}

func main() {
	fmt.Println(smallestEquivalentString("parker", "morris", "parser"))
}

/*
遍历 A B 字符串构造并查集，将值小的作为 root
遍历 S 字符串，对每个字母找到根 root 并拼接字符串
 */