package main
/*
971. 翻转二叉树以匹配先序遍历
给你一棵二叉树的根节点 root ，树中有 n 个节点，每个节点都有一个不同于其他节点
且处于 1 到 n 之间的值。
另给你一个由 n 个值组成的行程序列 voyage ，表示 预期 的二叉树 先序遍历 结果。
通过交换节点的左右子树，可以 翻转 该二叉树中的任意节点。例，翻转节点 1 的效果
如下：
请翻转 最少 的树中节点，使二叉树的 先序遍历 与预期的遍历行程 voyage 相匹配 。
如果可以，则返回 翻转的 所有节点的值的列表。你可以按任何顺序返回答案。如果不能，
则返回列表 [-1]。

示例 1
输入：root = [1,2], voyage = [2,1]
输出：[-1]
解释：翻转节点无法令先序遍历匹配预期行程。

示例 2：
输入：root = [1,2,3], voyage = [1,3,2]
输出：[1]
解释：交换节点 2 和 3 来翻转节点 1 ，先序遍历可以匹配预期行程。

示例 3：
输入：root = [1,2,3], voyage = [1,2,3]
输出：[]
解释：先序遍历已经匹配预期行程，所以不需要翻转节点。
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	position := new(int)
	output := make([]int, 0)
	*position = 0
	isMatch, output := helperFlip(root, voyage, position, output)
	if isMatch {
		if *position == len(voyage) {
			return output
		} else {
			return []int{-1}
		}
	} else {
		return []int{-1}
	}
}

func helperFlip(root *TreeNode, voyage []int, position *int, output []int) (bool, []int) {
	if root == nil {
		return true, output
	}
	if *position == len(voyage) {
		return false, output
	}
	if root.Val != voyage[*position] {
		return false, output
	}
	*position++
	leftIsMatch, output := helperFlip(root.Left, voyage, position, output)
	if leftIsMatch {
		return helperFlip(root.Right, voyage, position, output)
	} else {
		rightIsMatch, output := helperFlip(root.Right, voyage, position, output)
		if rightIsMatch {
			output = append(output, root.Val)
			return helperFlip(root.Left, voyage, position, output)
		} else {
			return false, output
		}
	}
}