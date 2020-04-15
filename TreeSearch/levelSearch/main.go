package levelSearch

import "github.com/yalbaba/DataStructures/queue"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelSearch(root *TreeNode) []int {
	var vals []int
	list := queue.Queue{}
	currentNode := root
	list.Push(currentNode)
	for {
		currentNode, _ := list.Pop()
		vals = append(vals, currentNode.(*TreeNode).Val)
		if currentNode.(*TreeNode).Left != nil {
			list.Push(currentNode)
		}
		if currentNode.(*TreeNode).Right != nil {
			list.Push(currentNode)
		}

		if list.IsEmpty() {
			break
		}
	}
	return vals
}
