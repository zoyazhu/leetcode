package main

import (
	"fmt"
	"sort"
)

/*
1178. 猜字谜
外国友人仿照中国字谜设计了一个英文版猜字谜小游戏，请你来猜猜看吧。
字谜的迷面 puzzle 按字符串形式给出，如果一个单词 word 符合下面两个条件，那么它就
可以算作谜底：
单词 word 中包含谜面 puzzle 的第一个字母。
单词 word 中的每一个字母都可以在谜面 puzzle 中找到。
例如，如果字谜的谜面是 "abcdefg"，那么可以作为谜底的单词有 "faced", "cabbage",
和 "baggage"；而 "beefed"（不含字母 "a"）以及 "based"（其中的 "s" 没有出现在
谜面中）都不能作为谜底。
返回一个答案数组 answer，数组中的每个元素 answer[i] 是在给出的单词列表 words
中可以作为字谜迷面 puzzles[i] 所对应的谜底的单词数目。

示例：
输入：
words = ["aaaa","asas","able","ability","actt","actor","access"],
puzzles = ["aboveyz","abrodyz","abslute","absoryz","actresz","gaswxyz"]
输出：[1,1,3,2,4,0]
解释：
1 个单词可以作为 "aboveyz" 的谜底 : "aaaa"
1 个单词可以作为 "abrodyz" 的谜底 : "aaaa"
3 个单词可以作为 "abslute" 的谜底 : "aaaa", "asas", "able"
2 个单词可以作为 "absoryz" 的谜底 : "aaaa", "asas"
4 个单词可以作为 "actresz" 的谜底 : "aaaa", "asas", "actt", "access"
没有单词可以作为 "gaswxyz" 的谜底，因为列表中的单词都不含字母 'g'。
 */

type treeNode struct {
	child [26]*treeNode
	count int
}

func findNumOfValidWords(words []string, puzzles []string) []int {
	root := &treeNode{}
	for _, word := range words {
		wordByte := []byte(word)
		sort.Slice(wordByte, func(i, j int) bool {
			return wordByte[i] < wordByte[j]
		})
		index := 0
		// 将word中的字符去重
		for _, c := range wordByte[1:] {
			if wordByte[index] != c {
				index++
				wordByte[index] = c
			}
		}
		wordByte = wordByte[:index+1]
		// 加入字典
		currentRoot := root
		for _, c := range wordByte {
			c = c - 'a'
			if currentRoot.child[c] == nil {
				currentRoot.child[c] = &treeNode{}
			}
			currentRoot = currentRoot.child[c]
		}
		currentRoot.count++
	}
	output := make([]int, len(puzzles))
	for i, puzzle := range puzzles {
		puzzleByte := []byte(puzzle)
		firstChar := puzzleByte[0]
		sort.Slice(puzzleByte, func(i, j int) bool {
			return puzzleByte[i] < puzzleByte[j]
		})
		// 在回溯的过程中枚举 pz 的所有子集并统计答案
		var findNode func(int, *treeNode) int
		findNode = func(position int, current *treeNode) int {
			if current == nil {
				return 0
			}
			if position == len(puzzle) {
				return current.count
			}
			result := findNode(position + 1, current.child[puzzleByte[position] - 'a'])
			// 当 pz[pos] 不为首字母时，可以不选择第 pos 个字母
			if puzzleByte[position] != firstChar {
				result = result + findNode(position + 1, current)
			}
			return result
		}
		output[i] = findNode(0, root)
	}
	return output
}

func main() {
	fmt.Println(findNumOfValidWords([]string{"aaaa","asas","able","ability","actt","actor","access"},
	[]string{"aboveyz","abrodyz","abslute","absoryz","actresz","gaswxyz"}))
}

/*
方法二：字典树
思路与算法
由于题目中规定word 和puzzle 均只包含小写字母，我们也可以考虑使用字典树来表
示需要的「数据结构」。由于方法一中已经详细介绍了每一步的做法，因此方法二中只
介绍与方法一不同的地方。
对于每一个word 对应的集合 S_w，我们将 S_w中的的字母按照字典序进行排序，组成
一个字符串，加入字典树中。与方法一中的哈希映射类似，我们需要统计每个字符串在
字典树中的出现次数。
对于每一个puzzle 对应的集合 S_p，我们枚举 S_p的子集，并将子集中的字母按照字典
序进行排序，组成一个字符串，在字典树中查询得到其出现次数。需要注意的是，由于 S_p
只是一个普通的集合，而不是二进制表示，因此我们可以使用回溯的方法，在枚举子集的同
时维护可以在字典树上进行查询的指针。详细的实现可以见下面的代码。

细节
在实际的实现中，我们无需显式地构造出加入字典树以及在字典树中查询需要的字符串，
可以考虑使用一些等价的简单数据结构（例如列表）表示字符串。
 */