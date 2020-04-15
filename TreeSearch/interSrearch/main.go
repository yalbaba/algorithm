package interSrearch

import "github.com/yalbaba/DataStructures/Stack"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//中序：从左开始找，第一次退回就打印
//递归中序
func InterOrderTree(root *TreeNode) []int {
	var vals []int
	if root != nil {
		InterOrderTree(root.Left)
		vals = append(vals, root.Val)
		InterOrderTree(root.Right)
	}
	return vals
}

//非递归中序
func InterOrderTree2(root *TreeNode) []int {
	var vals []int
	stack := Stack.New()
	stack.Push(root)
	currentNode, _ := stack.Top()
	for {
		if currentNode.(*TreeNode) != nil {
			stack.Push(currentNode)
			currentNode = currentNode.(*TreeNode).Left
		} else {
			topNode, _ := stack.Pop()
			vals = append(vals, topNode.(*TreeNode).Val)
			currentNode = topNode.(*TreeNode).Right
		}

		if (stack.Len() == 0) && currentNode == nil {
			break
		}
	}
	return vals
}
