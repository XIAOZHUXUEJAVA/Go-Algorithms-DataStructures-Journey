package trees

import (
	"fmt"
	"testing"
)

func TestBuildTree(t *testing.T) {
	preorder := []int{1, 2, 5, 4, 6, 7, 3, 8, 9}
	inorder := []int{5, 2, 6, 4, 7, 1, 8, 3, 9}
	root := buildTreeByPreIn(preorder, inorder)
	dispalyTreeByPre(root)
	fmt.Println()
	dispalyTreeByIn(root)
}
