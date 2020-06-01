package interSrearch

import "github.com/yalbaba/DataStructures/Stack"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//中序：从左开始找，第一次退回就打印
//递归中序
var res []int

func inorderTraversal(root *TreeNode) []int {
	res = make([]int, 0)
	dfs1(root)
	return res
}

func dfs1(root *TreeNode) {
	if root != nil {
		dfs1(root.Left)
		res = append(res, root.Val)
		dfs1(root.Right)
	}
}
//非递归中序
func inorderTraversal2(root *TreeNode) []int {
	res = make([]int, 0)
	dfs1(root)
	return res
}

func dfs2(root *TreeNode){
	stack := Stack.New()
	stack.Push(root)
	currentNode, _ := stack.Top()
	for {
		if currentNode != nil {
			stack.Push(currentNode)
			currentNode = currentNode.(*TreeNode).Left
		} else {
			topNode, _ := stack.Pop()
			res = append(res, topNode.(*TreeNode).Val)
			currentNode = topNode.(*TreeNode).Right
		}

		if (stack.Len() == 0) && currentNode == nil {
			break
		}
	}
}




