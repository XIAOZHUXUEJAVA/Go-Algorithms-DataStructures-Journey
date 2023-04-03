package trees

import (
	"fmt"
)

func dispalyTreeByPre(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Val, " ")
	dispalyTreeByPre(root.Left)
	dispalyTreeByPre(root.Right)
}

func dispalyTreeByIn(root *TreeNode) {
	if root == nil {
		return
	}
	dispalyTreeByIn(root.Left)
	fmt.Print(root.Val, " ")
	dispalyTreeByIn(root.Right)
}
