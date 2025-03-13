package Construct_Binary_Tree_from_Preorder_and_Inorder_Traversal

import (
	"testing"
)

// isEqual compares two binary trees for equality.
func isEqual(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	return t1.Val == t2.Val && isEqual(t1.Left, t2.Left) && isEqual(t1.Right, t2.Right)
}

var testCases = []struct {
	preorder []int
	inorder  []int
	expected *TreeNode
}{
	// Example 1:
	{
		preorder: []int{3, 9, 20, 15, 7},
		inorder:  []int{9, 3, 15, 20, 7},
		expected: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val: 15,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		},
	},
	// Example 2:
	{
		preorder: []int{-1},
		inorder:  []int{-1},
		expected: &TreeNode{
			Val: -1,
		},
	},
	// Left-skewed tree:
	// Tree structure:
	//       1
	//      /
	//     2
	//    /
	//   3
	//  /
	// 4
	{
		preorder: []int{1, 2, 3, 4},
		inorder:  []int{4, 3, 2, 1},
		expected: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
					},
				},
			},
		},
	},
	// Right-skewed tree:
	// Tree structure:
	// 1
	//  \
	//   2
	//    \
	//     3
	//      \
	//       4
	{
		preorder: []int{1, 2, 3, 4},
		inorder:  []int{1, 2, 3, 4},
		expected: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 4,
					},
				},
			},
		},
	},
	// A complete binary tree:
	// Tree structure:
	//         1
	//       /   \
	//      2     3
	//     / \   / \
	//    4   5 6   7
	{
		preorder: []int{1, 2, 4, 5, 3, 6, 7},
		inorder:  []int{4, 2, 5, 1, 6, 3, 7},
		expected: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
			Right: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		},
	},
}

// TestBuildTree tests the buildTree function with various cases.
func TestBuildTree(t *testing.T) {
	for _, tc := range testCases {
		result := buildTree(tc.preorder, tc.inorder)
		if !isEqual(result, tc.expected) {
			t.Errorf("buildTree(%v, %v) = %v, expected %v", tc.preorder, tc.inorder, result, tc.expected)
		}
	}
}

// TestBuildTree tests the buildTree function with various cases.
func TestBuildTree1(t *testing.T) {
	for _, tc := range testCases {
		result := buildTree1(tc.preorder, tc.inorder)
		if !isEqual(result, tc.expected) {
			t.Errorf("buildTree(%v, %v) = %v, expected %v", tc.preorder, tc.inorder, result, tc.expected)
		}
	}
}
