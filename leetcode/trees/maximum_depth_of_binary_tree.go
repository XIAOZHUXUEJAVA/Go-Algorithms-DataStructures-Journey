package trees

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMaxDepth := maxDepth(root.Left)
	rightMaxDepth := maxDepth(root.Right)
	res := maxInt(leftMaxDepth, rightMaxDepth) + 1
	return res
}
func maxInt(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
