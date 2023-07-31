package trees

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return checkLeftAndRight(root.Left, root.Right)
}

func checkLeftAndRight(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return checkLeftAndRight(left.Left, right.Right) && checkLeftAndRight(left.Right, right.Left)
}
