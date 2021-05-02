package main
/*
1080. 根到叶路径上的不足节点
给定一棵二叉树的根 root，请你考虑它所有 从根到叶的路径：从根到任何叶的路径。
（所谓一个叶子节点，就是一个没有子节点的节点）
假如通过节点 node 的每种可能的 “根-叶” 路径上值的总和全都小于给定的 limit，
则该节点被称之为「不足节点」，需要被删除。
请你删除所有不足节点，并返回生成的二叉树的根。

示例 1：
输入：root = [1,2,3,4,-99,-99,7,8,9,-99,-99,12,13,-99,14], limit = 1
输出：[1,2,3,4,null,null,7,8,9,null,14]

示例 2：
输入：root = [5,4,8,11,null,17,4,7,1,null,null,5,3], limit = 22
输出：[5,4,8,11,null,17,4,7,null,null,null,5]

示例 3：
输入：root = [5,-6,-6], limit = 0
输出：[]
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Left == nil && root.Right == nil {
		if root.Val < limit {
			return nil
		}
		return root
	}
	left := sufficientSubset(root.Left, limit - root.Val)
	right := sufficientSubset(root.Right, limit - root.Val)
	if left == nil && right == nil {
		return nil
	}
	root.Left = left
	root.Right = right
	return root
}

/*
从根结点出发，把值累加直到叶子节点。
如果 sum + root.Val < limit, 满足条件，直接把叶子节点置为nil，并返回false
分别判断左右两边的值，
如果一个跟的两个子节点都是false，那么经过这个根的sum都小于limit，把该节点置为nil
func sufficientSubset(root *TreeNode, limit int) *TreeNode {
    var l, r *TreeNode

    if root == nil {
        return nil
    }
    if root.Left == nil && root.Right == nil {
        if root.Val < limit {
            return nil
        }
        return root
    }

    l = sufficientSubset(root.Left, limit-root.Val)
    r = sufficientSubset(root.Right, limit-root.Val)

    if l==nil && r ==nil {
        return nil
    }


    root.Left = l
    root.Right = r
    return root
}
 */