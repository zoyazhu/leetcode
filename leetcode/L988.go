package main

/*
988. 从叶结点开始的最小字符串
给定一颗根结点为 root 的二叉树，树中的每一个结点都有一个从 0 到 25 的值，
分别代表字母 'a' 到 'z'：值 0 代表 'a'，值 1 代表 'b'，依此类推。
找出按字典序最小的字符串，该字符串从这棵树的一个叶结点开始，到根结点结束。
（小贴士：字符串中任何较短的前缀在字典序上都是较小的：例如，在字典序上
"ab" 比 "aba" 要小。叶结点是指没有子结点的结点。）

示例 1：
输入：[0,1,2,3,4,3,4]
输出："dba"

示例 2：
输入：[25,1,3,1,3,0,2]
输出："adz"

示例 3：
输入：[2,2,1,null,1,0,null,0]
输出："abc"
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func smallestFromLeaf(root *TreeNode) string {
	output := ""
	var traverse func(root *TreeNode, path []byte, output *string)
	traverse = func(root *TreeNode, path []byte, output *string) {
		if root == nil {
			return
		}
		path = append(path, byte(root.Val + 'a'))
		if root.Left == nil && root.Right == nil {
			number := reversePath(path)
			if *output == "" || *output > number {
				*output = number
			}
			return
		}
		traverse(root.Left, path, output)
		traverse(root.Right, path, output)
	}
	traverse(root, []byte{}, &output)
	return output
}

func reversePath(path []byte) string {
	output := ""
	for _, v := range string(path) {
		output = string(v) + output
	}
	return output
}