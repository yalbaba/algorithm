package main

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func isSameTree(root1, root2 *TreeNode) bool {
	if root1 != nil && root2 == nil {
		return false
	}
	if root2 != nil && root1 == nil {
		return false
	}
	if root1 == nil && root2 == nil {
		return true
	}
	if root1.val != root2.val {
		return false
	}
	return isSameTree(root1.left, root2.left) && isSameTree(root1.right, root2.right)
}

func main() {

}
