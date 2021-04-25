package main
/*
865. 具有所有最深节点的最小子树
给定一个根为 root 的二叉树，每个节点的深度是 该节点到根的最短距离 。
如果一个节点在 整个树 的任意节点之间具有最大的深度，则该节点是 最深的 。
一个节点的 子树 是该节点加上它的所有后代的集合。
返回能满足 以该节点为根的子树中包含所有最深的节点 这一条件的具有最大深度的节点。

示例 1：
输入：root = [3,5,1,6,2,0,8,null,null,7,4]
输出：[2,7,4]
解释：
我们返回值为 2 的节点，在图中用黄色标记。
在图中用蓝色标记的是树的最深的节点。
注意，节点 5、3 和 2 包含树中最深的节点，但节点 2 的子树最小，因此我们返回它。

示例 2：
输入：root = [1]
输出：[1]
解释：根节点是树中最深的节点。

示例 3：
输入：root = [0,1,3,null,2]
输出：[2]
解释：树中最深的节点为 2 ，有效子树为节点 2、1 和 0 的子树，但节点 2 的子树最小。
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	_, output := subTreeDepth(root)
	return output
}

func subTreeDepth(root *TreeNode) (int, *TreeNode) {
	if root == nil {
		return 0, root
	}
	leftDepth, left := subTreeDepth(root.Left)
	rightDepth, right := subTreeDepth(root.Right)
	if leftDepth == rightDepth {
		return leftDepth + 1, root
	}
	if leftDepth > rightDepth {
		return leftDepth + 1, left
	}
	return rightDepth + 1, right
}

/*
解题思路
如果左右子树的深度一样，那么返回的是当前父节点；否则谁的深度更深就把
谁给返回去。逻辑就是这个样子

 */