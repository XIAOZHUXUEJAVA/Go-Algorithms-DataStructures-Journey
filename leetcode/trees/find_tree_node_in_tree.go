package trees

func searchNodeInTree(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == target {
		return root
	}
	left := searchNodeInTree(root.Left, target)
	if left != nil {
		return left
	}
	right := searchNodeInTree(root.Right, target)
	if right != nil {
		return right
	}
	return nil
}
