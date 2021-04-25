package main
/*
951. 翻转等价二叉树
我们可以为二叉树 T 定义一个翻转操作，如下所示：选择任意节点，然后交换它的左子树
和右子树。
只要经过一定次数的翻转操作后，能使 X 等于 Y，我们就称二叉树 X 翻转等价于二叉树 Y。
编写一个判断两个二叉树是否是翻转等价的函数。这些树由根节点 root1 和 root2 给出。

示例：
输入：root1 = [1,2,3,4,5,6,null,null,null,7,8],
root2 = [1,3,2,null,6,4,5,null,null,null,null,8,7]
输出：true
解释：我们翻转值为 1，3 以及 5 的三个节点。
Flipped Trees Diagram
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	} else if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val == root2.Val {
		return flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left) ||
			flipEquiv(root1.Right, root2.Right) && flipEquiv(root1.Left, root2.Left)
	}
	return false
}

