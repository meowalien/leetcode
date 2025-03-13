package Construct_Binary_Tree_from_Preorder_and_Inorder_Traversal

/*
Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree and inorder is the inorder traversal of the same tree, construct and return the binary tree.

Example 1:

Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]
Example 2:

Input: preorder = [-1], inorder = [-1]
Output: [-1]

Constraints:

1 <= preorder.length <= 3000
inorder.length == preorder.length
-3000 <= preorder[i], inorder[i] <= 3000
preorder and inorder consist of unique values.
Each value of inorder also appears in preorder.
preorder is guaranteed to be the preorder traversal of the tree.
inorder is guaranteed to be the inorder traversal of the tree.
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	rootVal := preorder[0]
	cutPoint := indexOf(inorder, rootVal)
	if cutPoint == -1 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{
			Val: rootVal,
		}
	}

	var rightTree *TreeNode
	if cutPoint+1 < len(inorder) {
		rightTree = buildTree(preorder[cutPoint+1:], inorder[cutPoint+1:])
	} else {
		rightTree = nil
	}

	return &TreeNode{
		Val:   rootVal,
		Left:  buildTree(preorder[1:], inorder[:cutPoint]),
		Right: rightTree,
	}
}

func indexOf(inorder []int, val int) int {
	for i, v := range inorder {
		if v == val {
			return i
		}
	}
	return -1
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
