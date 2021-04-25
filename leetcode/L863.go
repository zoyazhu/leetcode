package main
/*
863. 二叉树中所有距离为 K 的结点
给定一个二叉树（具有根结点 root）， 一个目标结点 target ，和一个整数值 K 。
返回到目标结点 target 距离为 K 的所有结点的值的列表。 答案可以以任何顺序返回。

示例 1：
输入：root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, K = 2
输出：[7,4,1]
解释：
所求结点为与目标结点（值为 5）距离为 2 的结点，
值分别为 7，4，以及 1
注意，输入的 "root" 和 "target" 实际上是树上的结点。
上面的输入仅仅是对这些对象进行了序列化描述。

提示：
给定的树是非空的。
树上的每个结点都具有唯一的值 0 <= node.val <= 500 。
目标结点 target 是树上的结点。
0 <= K <= 1000.
 */

// Definition for a binary tree node.
type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	output := make([]int, 0)
	if root == nil {
		return output
	}
	// 存储root到target的父节点
	parent := make([]*TreeNode, 0, 500)
	var parentPath func(root, target *TreeNode) bool
	parentPath = func(root, target *TreeNode) bool {
		if root == nil {
			return false
		}
		if root == target {
			return true
		}
		parent = append(parent, root)
		if parentPath(root.Left, target) || parentPath(root.Right, target) {
			return true
		}
		parent = parent[:len(parent) - 1]
		return false
	}
	parentPath(root, target)
	// 从target的子树下得到所有节点
	output = append(output, findPath(target, K)...)
	// 从target的父节点得到节点
	depth := 1
	preNode := target
	for len(parent) > 0 && depth <= K {
		lastParent := parent[len(parent) - 1]
		parent = parent[:len(parent) - 1]
		if K == depth {
			output = append(output, lastParent.Val)
		} else if preNode == lastParent.Left {
			output = append(output, findPath(lastParent.Right, K - depth - 1)...)
		} else {
			output = append(output, findPath(lastParent.Left, K - depth - 1)...)
		}
		depth++
		preNode = lastParent
	}
	return output
}

func findPath(root *TreeNode, K int) []int {
	output := make([]int, 0)
	if root == nil {
		return output
	}
	if K == 0 {
		output = append(output, root.Val)
		return output
	}
	left := findPath(root.Left, K - 1)
	right := findPath(root.Right, K - 1)
	output = append(output, left...)
	output = append(output, right...)
	return output
}

