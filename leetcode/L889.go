package main
/*
889. 根据前序和后序遍历构造二叉树
返回与给定的前序和后序遍历匹配的任何二叉树。
 pre 和 post 遍历中的值是不同的正整数。

示例：
输入：pre = [1,2,4,5,3,6,7], post = [4,5,2,6,7,3,1]
输出：[1,2,3,4,5,6,7]
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructFromPrePost(pre []int, post []int) *TreeNode {
	if len(post) == 0 {
		return nil
	}
	if len(post) == 1 {
		return &TreeNode{Val: post[0]}
	}
	root := new(TreeNode)
	root.Val = pre[0]
	rightChild := post[len(post) - 2]
	index := 0
	for k, v := range pre {
		if v == rightChild {
			index = k
			break
		}
	}
	rightTree := len(pre) - index
	leftTree := len(post) - 1 - rightTree

	preLeft := pre[1:leftTree + 1]
	preRight := pre[index:]

	postLeft := post[:leftTree]
	postRight := post[leftTree:len(post) - 1]
	root.Left = constructFromPrePost(preLeft, postLeft)
	root.Right = constructFromPrePost(preRight, postRight)
	return root
}

/*
代码思路:
1.递归处理。
2.终止条件:
遍历的数组为1时则返回这个节点；
遍历的数组为空时则返回一个nil；
3.区间划分：
设定一个哨兵索引，方便进行计算
4.递归进行：
左子树，右子树分开递归，递归结果分别返回为左子节点和右子节点。
 */