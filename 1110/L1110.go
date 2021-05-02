package main
/*
1110. 删点成林
给出二叉树的根节点 root，树上每个节点都有一个不同的值。
如果节点值在 to_delete 中出现，我们就把该节点从树上删去，最后得到一个森林（一些不相交的树构成的集合）。
返回森林中的每棵树。你可以按任意顺序组织答案。

示例：
输入：root = [1,2,3,4,5,6,7], to_delete = [3,5]
输出：[[1,2,null,4],[6],[7]]
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	output := []*TreeNode{}
	dummy := &TreeNode{0, root, nil}
	deleteMap := make(map[int]bool)
	deleteMap[0] = true
	for _, i := range to_delete {
		deleteMap[i] = true
	}
	var dfs func(root * TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root != nil {
			if dfs(root.Left) {
				root.Left = nil
			}
			if dfs(root.Right) {
				root.Right = nil
			}
			if deleteMap[root.Val] {
				for _, i := range []*TreeNode{root.Left, root.Right} {
					if i != nil {
						output = append(output, i)
					}
				}
				return true
			}
		}
		return false
	}
	dfs(dummy)
	return output
}

/*
思路：假如说当前节点root是需要删除的 但是root左右子树中还有需要删除的节点roots那么优先删除roots然后再删除root是比较直观明了的 也就是说需要使用到后序遍历
注意：基于树的性质其他遍历方式也是可以的 本人菜鸡 主要的语句都添加了注释 一下代码仅供参考 还请大家多多指教

func delNodes(root *TreeNode,to_delete []int) (ans []*TreeNode) {
	None:=(*TreeNode)(nil)
	dummy:=&TreeNode{0,root,None} //一个哑节点 方便后续的统一处理(比如说root最后可能是要放进ans中去的)
	delete:=map[int]bool{0:true} //[0]就表示要删除dummy节点 []int->map 简单地模拟一下set 实现快速查询
	for _,i:=range to_delete {
		delete[i]=true
	}
	var dfs func (root *TreeNode) bool
	dfs=func(root *TreeNode) bool {
		if root!=nil {
			if dfs(root.Left) {
				root.Left=None
			}
			if dfs(root.Right) {
				root.Right=None
			}
			if delete[root.Val] {
				for _,i:=range []*TreeNode{root.Left,root.Right} {
					if i!=None { //None是不需要的
						ans=append(ans,i)
					}
				}
				return true //true表示当前节点是要删掉的 用来提醒该节点的双亲节点去None
			}
		}
		return false
	}
	dfs(dummy)
	return
}
 */
