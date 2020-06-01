package main

import "github.com/yalbaba/DataStructures/Stack"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


//前序：从左开始找，找到一个打印一个

var res []int
//递归前序遍历
func PreOrderTree(root *TreeNode) []int {
	res = []int{}
	Pre1(root)
	return res
}

func Pre1(root *TreeNode) {
	if root != nil {
		res = append(res, root.Val)
		//	遍历左节点
		Pre1(root.Left)
		//	遍历右节点
		Pre1(root.Right)
	}
}

//非递归前序遍历
func Pre2(root *TreeNode){
	stack := Stack.New()
	stack.Push(root)
	currentNode, _ := stack.Top()
	for {
		if currentNode != nil {
			res = append(res, currentNode.(*TreeNode).Val)
			//入栈，找左节点
			stack.Push(currentNode)
			currentNode = currentNode.(*TreeNode).Left
		} else {
			//出栈顶，找右节点
			topNode, _ := stack.Pop()
			currentNode = topNode.(*TreeNode).Right
		}

		if (stack.Len() == 0) && currentNode == nil {
			break
		}
	}
}
