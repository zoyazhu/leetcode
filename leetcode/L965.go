package main
/*
965. 单值二叉树
如果二叉树每个节点都具有相同的值，那么该二叉树就是单值二叉树。
只有给定的树是单值二叉树时，才返回 true；否则返回 false。

示例 1：
输入：[1,1,1,1,1,null,1]
输出：true

示例 2：
输入：[2,2,2,5,2]
输出：false
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && root.Left.Val != root.Val {
		return false
	}
	if root.Right != nil && root.Right.Val != root.Val {
		return false
	}
	output := true
	output = output && isUnivalTree(root.Left) && isUnivalTree(root.Right)
	return output
}
