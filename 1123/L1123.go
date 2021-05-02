package main
/*
1123. 最深叶节点的最近公共祖先
给你一个有根节点的二叉树，找到它最深的叶节点的最近公共祖先。
回想一下：
叶节点 是二叉树中没有子节点的节点
树的根节点的 深度 为 0，如果某一节点的深度为 d，那它的子节点的深度就是 d+1
如果我们假定 A 是一组节点 S 的 最近公共祖先，S 中的每个节点都在以 A 为根节点的子树中，且 A 的深度达到此条件下可能的最大值。
注意：本题与力扣 865 重复：https://leetcode-cn.com/problems/smallest-subtree-with-all-the-deepest-nodes/

示例 1：
输入：root = [3,5,1,6,2,0,8,null,null,7,4]
输出：[2,7,4]
解释：
我们返回值为 2 的节点，在图中用黄色标记。
在图中用蓝色标记的是树的最深的节点。
注意，节点 6、0 和 8 也是叶节点，但是它们的深度是 2 ，而节点 7 和 4 的深度是 3 。

示例 2：
输入：root = [1]
输出：[1]
解释：根节点是树中最深的节点，它是它本身的最近公共祖先。

示例 3：
输入：root = [0,1,3,null,2]
输出：[2]
解释：树中最深的叶节点是 2 ，最近公共祖先是它自己。
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	var dfs func(root *TreeNode, depth int) (*TreeNode, int)
	dfs = func(root *TreeNode, depth int) (*TreeNode, int) {
		if root == nil {
			return root, depth
		}
		left, leftDepth := dfs(root.Left, depth+1)
		right, rightDepth := dfs(root.Right, depth+1)
		if leftDepth == rightDepth {
			return root, leftDepth
		}
		if leftDepth > rightDepth {
			return left, leftDepth
		}
		return right, rightDepth
	}
	output, _ := dfs(root, 0)
	return output
}

/*
代码

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	ret, _ := dfs(root, 0)
	return ret
}

这里返回的是以root为根节点的树的最深深度叶子节点的LCA，以及最大深度
func dfs(root *TreeNode, fromDepth int) (leafLcaNode *TreeNode, finalDepth int) {
	if nil == root {
		return root, fromDepth
	}
	lNode, lDep := dfs(root.Left, fromDepth + 1)
	rNode, rDep := dfs(root.Right, fromDepth + 1)
	if lDep == rDep {
		return root, lDep
	}
	if lDep > rDep {
		return lNode, lDep
	}
	return rNode, rDep
}
 */
