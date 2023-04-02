package trees

func lowestCommonAncestorBST(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if p.Val > root.Val && q.Val > root.Val {
			root = root.Right
		} else if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
		} else {
			break
		}
	}
	return root
}
