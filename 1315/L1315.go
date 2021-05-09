package main
/*
1315. 祖父节点值为偶数的节点和
给你一棵二叉树，请你返回满足以下条件的所有节点的值之和：
该节点的祖父节点的值为偶数。（一个节点的祖父节点是指该节点的父节点的父节点。）
如果不存在祖父节点值为偶数的节点，那么返回 0 。

示例：
输入：root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
输出：18
解释：图中红色节点的祖父节点的值为偶数，蓝色节点为这些红色节点的祖父节点。
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumEvenGrandparent(root *TreeNode) int {
	output := 0
	var dfs func(root *TreeNode, parent, grandparent bool)
	dfs = func(root *TreeNode, parent, grandparent bool) {
		if root != nil {
			if grandparent {
				output = output + root.Val
			}
			dfs(root.Left, root.Val%2 == 0, parent)
			dfs(root.Right, root.Val%2 == 0, parent)
		}
	}
	dfs(root, false, false)
	return output
}

