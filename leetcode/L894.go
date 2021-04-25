package main
/*
894. 所有可能的满二叉树
满二叉树是一类二叉树，其中每个结点恰好有 0 或 2 个子结点。
返回包含 N 个结点的所有可能满二叉树的列表。 答案的每个元素都是一个可能树的根结点。
答案中每个树的每个结点都必须有 node.val=0。
你可以按任何顺序返回树的最终列表。

示例：
输入：7
输出：[[0,0,0,null,null,0,0,null,null,0,0],[0,0,0,null,null,0,0,0,0]
,[0,0,0,0,0,0,0],[0,0,0,0,0,null,null,null,null,0,0],
[0,0,0,0,0,null,null,0,0]]
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func allPossibleFBT(n int) []*TreeNode {
	if n % 2 == 0 {
		return nil
	}
	if n == 1 {
		return []*TreeNode{new(TreeNode)}
	}
	output := make([]*TreeNode, 0)
	for i := 1; i < n - 1; i = i + 2 {
		left := allPossibleFBT(i)
		right := allPossibleFBT(n - i - 1)
		for _, l := range left {
			for _, r := range right {
				node := new(TreeNode)
				node.Left = l
				node.Right = r
				output = append(output, node)
			}
		}
	}
	return output
}

/*
解题思路
N为偶数时无解，二维数组存树，本质四重循环，多节点的二叉树都是由少节点的二叉树
组合而成，返回的二叉树数组实质在共用子树。
代码里d[n]为节点数是n的二叉树数组，i则该二叉树代表左侧分配的节点数，n - i - 1
则为右侧分配的节点数，再遍历组合左右节点数对应的二叉树数组即可得到n个节点时的二
叉树数组。
 */