package main
/*
872. 叶子相似的树
请考虑一棵二叉树上所有的叶子，这些叶子的值按从左到右的顺序排列形成一个 叶值序列 。
举个例子，如上图所示，给定一棵叶值序列为 (6, 7, 4, 9, 8) 的树。
如果有两棵二叉树的叶值序列是相同，那么我们就认为它们是 叶相似 的。
如果给定的两个头结点分别为 root1 和 root2 的树是叶相似的，则返回 true；
否则返回 false 。

示例 1：
输入：root1 = [3,5,1,6,2,9,8,null,null,7,4], root2 = [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]
输出：true

示例 2：
输入：root1 = [1], root2 = [1]
输出：true

示例 3：
输入：root1 = [1], root2 = [2]
输出：false

示例 4：
输入：root1 = [1,2], root2 = [2,2]
输出：true

示例 5：
输入：root1 = [1,2,3], root2 = [1,3,2]
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
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leafList1 := make([]int, 0)
	leafList2 := make([]int, 0)
	dfs872(root1, &leafList1)
	dfs872(root2, &leafList2)
	if len(leafList1) != len(leafList2) {
		return false
	}
	for i := 0; i < len(leafList2); i++ {
		if leafList2[i] != leafList1[i] {
			return false
		}
	}
	return true
}

func dfs872(root *TreeNode, leaf *[]int) {
	if root != nil {
		dfs872(root.Left, leaf)
		if root.Left == nil && root.Right == nil {
			*leaf = append(*leaf, root.Val)
		}
		dfs872(root.Right, leaf)
	}
}

/*
方法：深度优先搜索
思路和算法
首先，让我们找出给定的两个树的叶值序列。之后，我们可以比较它们，看看它们是否相等。
要找出树的叶值序列，我们可以使用深度优先搜索。如果结点是叶子，那么 dfs 函数会写入
结点的值，然后递归地探索每个子结点。这可以保证按从左到右的顺序访问每片叶子，因为
在右孩子结点之前完全探索了左孩子结点。
 */