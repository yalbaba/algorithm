package afterSearch

import "github.com/yalbaba/DataStructures/Stack"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//后序递归
func AfterSearch(root *TreeNode) []int {
	var vals []int
	if root != nil {
		AfterSearch(root.Left)
		AfterSearch(root.Right)
		vals = append(vals, root.Val)
	}
	return vals
}

func AfterSearch2(root *TreeNode) []int {
	stack := Stack.New()
	var vals []int
	stack.Push(root)
	currentNode, _ := stack.Top()
	var temp *TreeNode
	for {
		if currentNode != nil {
			stack.Push(currentNode)
			currentNode = currentNode.(*TreeNode).Left
		} else {
			currentNode, _ = stack.Top()
			if currentNode.(*TreeNode).Right != nil && temp != currentNode.(*TreeNode).Right {
				currentNode = currentNode.(*TreeNode).Right
			} else {
				currentNode, _ = stack.Pop()
				vals = append(vals, currentNode.(*TreeNode).Val)
				temp = currentNode.(*TreeNode)
				currentNode = nil
			}
		}

		if (stack.Len() == 0) && currentNode == nil {
			break
		}
	}
	return vals
}
