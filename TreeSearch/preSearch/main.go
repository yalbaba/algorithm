package main

import "github.com/yalbaba/DataStructures/Stack"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

//前序：从左开始找，找到一个打印一个

//递归前序遍历
func PreOrderTree(root *TreeNode) []int {
	var vals []int
	if root != nil {
		vals = append(vals, root.Val)
		//	遍历左节点
		PreOrderTree(root.Left)
		//	遍历右节点
		PreOrderTree(root.Right)
	}
	return vals
}

//非递归前序遍历
func PreOrderTree2(root *TreeNode) []int {
	var vals []int
	stack := Stack.New()
	stack.Push(root)
	currentNode, _ := stack.Top()
	for {
		if currentNode != nil {
			vals = append(vals, currentNode.(*TreeNode).Val)
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
	return vals
}
