package main
/*
897. 递增顺序查找树
给你一个树，请你 按中序遍历 重新排列树，使树中最左边的结点现在是树的根，
并且每个结点没有左子结点，只有一个右子结点。

示例 ：
输入：[5,3,6,2,4,null,8,1,null,null,null,7,9]

       5
      / \
    3    6
   / \    \
  2   4    8
 /        / \
1        7   9
输出：[1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]

 1
  \
   2
    \
     3
      \
       4
        \
         5
          \
           6
            \
             7
              \
               8
                \
                 9
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
	var dfs func(root *TreeNode, nums *[]int)
	nums := make([]int, 0)
	dfs = func(root *TreeNode, nums *[]int) {
		if root == nil {
			return
		}
		dfs(root.Left, nums)
		*nums = append(*nums, root.Val)
		dfs(root.Right, nums)
	}
	dfs(root, &nums)
	output := &TreeNode{
		Val: nums[0],
		Left: nil,
		Right: nil,
	}
	current := output
	for _, v := range nums[1:] {
		nextNode := &TreeNode{
			Val: v,
			Left: nil,
			Right: nil,
		}
		current.Right = nextNode
		current = nextNode
	}
	return output
}

