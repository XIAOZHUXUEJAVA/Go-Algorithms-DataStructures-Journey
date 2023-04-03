package trees

import (
	"fmt"
	"reflect"
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

func TestBinaryTree(t *testing.T) {
	preorder := []int{1, 2, 5, 4, 6, 7, 3, 8, 9}
	inorder := []int{5, 2, 6, 4, 7, 1, 8, 3, 9}
	root := buildTreeByPreIn(preorder, inorder)

	t.Run("Test maxDepth()", func(t *testing.T) {
		got := maxDepth(root)
		expected := 3
		// expected := 4
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got: %d, want: %d", got, expected)
		}
	})

}

func TestBinarySearchTree(t *testing.T) {
	preorder := []int{4, 2, 1, 3, 7}
	inorder := []int{1, 2, 3, 4, 7}
	root := buildTreeByPreIn(preorder, inorder)

	t.Run("Test lowestCommonAncestorBST", func(t *testing.T) {
		p := searchNodeInTree(root, 2)
		q := searchNodeInTree(root, 7)
		expected := searchNodeInTree(root, 4)
		// 还得找到相应的二叉树
		got := lowestCommonAncestorBST(root, p, q)
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got: %d, expected: %d", got.Val, expected.Val)
		}
	})
}
