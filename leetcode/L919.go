package main
/*
919. 完全二叉树插入器
完全二叉树是每一层（除最后一层外）都是完全填充（即，节点数达到最大）的，并
且所有的节点都尽可能地集中在左侧。
设计一个用完全二叉树初始化的数据结构 CBTInserter，它支持以下几种操作：
CBTInserter(TreeNode root) 使用头节点为 root 的给定树初始化该数据结构；
CBTInserter.insert(int v)  向树中插入一个新节点，节点类型为 TreeNode，值为 v 。使树保持完全二叉树的状态，并返回插入的新节点的父节点的值；
CBTInserter.get_root() 将返回树的头节点。
示例 1：
输入：inputs = ["CBTInserter","insert","get_root"], inputs = [[[1]],[2],[]]
输出：[null,1,[1,2]]

示例 2：
输入：inputs = ["CBTInserter","insert","insert","get_root"], inputs = [[[1,2,3,4,5,6]],[7],[8],[]]
输出：[null,3,4,[1,2,3,4,5,6,7,8]]
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type CBTInserter struct {
	tree []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	this := &CBTInserter{[]*TreeNode{root}}
	for i := 0; i < len(this.tree); i++ {
		if this.tree[i].Left != nil {
			this.tree = append(this.tree, this.tree[i].Left)
		}
		if this.tree[i].Right != nil {
			this.tree = append(this.tree, this.tree[i].Right)
		}
	}
	return *this
}

func (this *CBTInserter) Insert(v int) int {
	length := len(this.tree)
	father := this.tree[(length - 1) >> 1]
	this.tree = append(this.tree, &TreeNode{Val: v})
	if length % 2 == 1 {
		father.Left = this.tree[length]
	} else {
		father.Right = this.tree[length]
	}
	return father.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.tree[0]
}


/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(v);
 * param_2 := obj.Get_root();
 */