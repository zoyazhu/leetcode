package main

import _ "go/build"

/*
430. 扁平化多级双向链表
多级双向链表中，除了指向下一个节点和前一个节点指针之外，它还有一个子链表指针，
可能指向单独的双向链表。这些子列表也可能会有一个或多个自己的子项，依此类推，
生成多级数据结构，如下面的示例所示。
给你位于列表第一级的头节点，请你扁平化列表，使所有结点出现在单级双链表中。

示例 1：
输入：head = [1,2,3,4,5,6,null,null,null,7,8,9,10,null,null,11,12]
输出：[1,2,3,7,8,11,12,9,10,4,5,6]
解释：
输入的多级列表如下图所示：
扁平化后的链表如下图：

示例 2：
输入：head = [1,2,null,3]
输出：[1,3,2]
解释：
输入的多级列表如下图所示：
  1---2---NULL
  |
  3---NULL

示例 3：
输入：head = []
输出：[]
 */
//Definition for a Node.
  type Node struct {
      Val int
      Prev *Node
      Next *Node
      Child *Node
  }

func flatten(root *Node) *Node {
	list := make([]*Node, 0)
	var dfs func(root *Node, list *[]*Node)
	dfs = func(root *Node, list *[]*Node) {
		if root == nil {
			return
		}
		*list = append(*list, root)
		dfs(root.Child, list)
		dfs(root.Next, list)
	}
	dfs(root, &list)
	length := len(list)
	for k, v := range list {
		if k + 1 < length {
			v.Next = list[k + 1]
		}
		if k > 0 {
			v.Prev = list[k - 1]
		}
		v.Child = nil
	}
	return root
}