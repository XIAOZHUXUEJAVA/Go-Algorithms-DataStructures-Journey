package trees

func buildTreeByPreIn(preorder, inorder []int) *TreeNode {

	if len(preorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	index := 0
	for i := 0; i <= len(inorder)-1; i++ {
		if rootVal == inorder[i] {
			index = i
			break
		}
	}
	root := &TreeNode{rootVal, nil, nil}

	// 1 2 5 4 6 7 3 8 9
	// 5 2 6 4 7 1 8 3 9
	// left 1:6 len(preorder[:index]) + 1
	root.Left = buildTreeByPreIn(preorder[1:len(preorder[:index])+1], inorder[:index])
	root.Right = buildTreeByPreIn(preorder[len(preorder[:index])+1:], inorder[index+1:])
	return root
}
