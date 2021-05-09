package main
/*
1302. 层数最深叶子节点的和
给你一棵二叉树的根节点 root ，请你返回 层数最深的叶子节点的和 。

示例 1：
输入：root = [1,2,3,4,5,null,6,7,null,null,null,null,8]
输出：15
示例 2：

输入：root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
输出：19
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
	var dfs func(root *TreeNode, depth int) (int, int)
	dfs = func(root *TreeNode, depth int) (int, int) {
		if root == nil {
			return 0, 0
		}
		leftMax, rightMax, leftDepth, rightDepth := 0, 0, 0, 0
		depth++
		if root.Left != nil {
			leftDepth, leftMax = dfs(root.Left, depth)
		}
		if root.Right != nil {
			rightDepth, rightMax = dfs(root.Right, depth)
		}
		if root.Left == nil && root.Right == nil {
			return depth, root.Val
		}
		if leftDepth > rightDepth {
			return leftDepth, leftMax
		} else if rightDepth == leftDepth {
			return rightDepth, rightMax + leftMax
		} else {
			return rightDepth, rightMax
		}
	}
	_, output := dfs(root, 0)
	return output
}