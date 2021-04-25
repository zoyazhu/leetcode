package main

import "fmt"

/*
839. 相似字符串组
如果交换字符串 X 中的两个不同位置的字母，使得它和字符串 Y 相等，
那么称 X 和 Y两个字符串相似。如果这两个字符串本身是相等的，那它们也是相似的。
例如，"tars" 和 "rats" 是相似的 (交换 0 与 2 的位置)；
"rats" 和 "arts" 也是相似的，但是 "star" 不与 "tars"，"rats"，或 "arts" 相似。
总之，它们通过相似性形成了两个关联组：{"tars", "rats", "arts"} 和 {"star"}。
注意，"tars" 和 "arts" 是在同一组中，即使它们并不相似。形式上，对每个
组而言，要确定一个单词在组中，只需要这个词和该组中至少一个单词相似。

给你一个字符串列表 strs。列表中的每个字符串都是 strs 中其它所有字符串
的一个字母异位词。请问 strs 中有多少个相似字符串组？

示例 1：
输入：strs = ["tars","rats","arts","star"]
输出：2

示例 2：
输入：strs = ["omv","ovm"]
输出：1
 */

func numSimilarGroups(strs []string) int {
	length := len(strs)
	uf := newUnionFind(length)
	for i, s := range strs {
		for j := i + 1; j < length; j++ {
			if !uf.isInSameSet(i, j) && uf.isSimilar(s, strs[j]) {
				uf.union(i, j)
			}
		}
	}
	return uf.count
}

type unionFind struct {
	parent []int
	size []int
	count int
}

func newUnionFind(n int) *unionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return &unionFind{parent: parent, size: size, count: n}
}

func (u *unionFind) isInSameSet(i, j int) bool {
	return u.find(i) == u.find(j)
}

func (u *unionFind) isSimilar(s, t string) bool {
	different := 0
	for i := range s {
		if s[i] != t[i] {
			different++
			if different > 2 {
				return false
			}
		}
	}
	return true
}

func (u *unionFind) union(i, j int) {
	parentI := u.find(i)
	parentJ := u.find(j)
	if parentI == parentJ {
		return
	}
	if u.size[parentJ] > u.size[parentI] {
		parentI, parentJ = parentJ, parentI
	}
	u.size[parentI] = u.size[parentI] + u.size[parentJ]
	u.parent[parentJ] = parentI
	u.count--
}

func (u *unionFind) find(i int) int {
	if u.parent[i] != i {
		u.parent[i] = u.find(u.parent[i])
	}
	return u.parent[i]
}

func main() {
	fmt.Println(numSimilarGroups([]string{"tars","rats","arts","star"}))
}

/*
方法一：并查集
思路及解法
我们把每一个字符串看作点，字符串之间是否相似看作边，那么可以发现
本题询问的是给定的图中有多少连通分量。于是可以想到使用并查集维护
节点间的连通性。

我们枚举给定序列中的任意一对字符串，检查其是否具有相似性，如果相似，
那么我们就将这对字符串相连。
在实际代码中，我们可以首先判断当前这对字符串是否已经连通，
如果没有连通，我们再检查它们是否具有相似性，可以优化一定的时间复杂度的常数。

 */
