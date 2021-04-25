package main

import "strconv"

/*
428. 序列化和反序列化 N 叉树
序列化是指将一个数据结构转化为位序列的过程，因此可以将其存储在文件中或
内存缓冲区中，以便稍后在相同或不同的计算机环境中恢复结构。
设计一个序列化和反序列化 N 叉树的算法。一个 N 叉树是指每个节点都有不超
过 N 个孩子节点的有根树。序列化 / 反序列化算法的算法实现没有限制。你只
需要保证 N 叉树可以被序列化为一个字符串并且该字符串可以被反序列化成原树
结构即可。
例如，你需要序列化下面的 3-叉 树。
为 [1 [3[5 6] 2 4]]。你不需要以这种形式完成，你可以自己创造和实现不同
的方法。
或者，您可以遵循 LeetCode 的层序遍历序列化格式，其中每组孩子节点由空值分隔。

例如，上面的树可以序列化为 [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,
null,null,11,null,12,null,13,null,null,14]
你不一定要遵循以上建议的格式，有很多不同的格式，所以请发挥创造力，想出不
同的方法来完成本题。

// Definition for a Node.
 */
type Node1 struct {
     Val int
     Children []*Node1
}

type Codec struct {
}

//func Constructor() *Codec {
//	return &Codec{}
//}

func (this *Codec) serialize(root *Node1) string {
	if root == nil {
		return "[]"
	}
	output := strconv.Itoa(root.Val) + ","
	if root.Children == nil {
		return output
	}
	output = output + "["
	for _, c := range root.Children {
		output = output + this.serialize(c)
	}
	return output + "]"
}

func (this *Codec) deserialize(data string) *Node1 {
	if data == "[]" {
		return nil
	}
	root := &Node1{}
	output := root
	sum := 0
	stack := make([]*Node1, 0)
	currentNode := &Node1{}
	for _, c := range data {
		switch true {
		case c >= 48 && c <= 57:
			sum = sum * 10 + int(c - 48)
		case c == ',':
			currentNode = &Node1{Val: sum}
			root.Children = append(root.Children, currentNode)
			sum = 0
		case c == '[':
			stack = append(stack, root)
			root = currentNode
			sum = 0
		case c == ']':
			root = stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
		}
	}
	return output.Children[0]
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */