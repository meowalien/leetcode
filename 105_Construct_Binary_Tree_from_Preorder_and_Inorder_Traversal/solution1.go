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
func buildTree1(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	// 建立一個 map，用於快速查找節點值在 inorder 中的索引位置
	inMap := make(map[int]int)
	for i, val := range inorder {
		inMap[val] = i
	}

	// 用於追蹤目前在 preorder 陣列的位置
	preIndex := 0

	// 閉包函式：負責根據左右邊界 (l, r) 建構子樹
	var helper func(l, r int) *TreeNode
	helper = func(l, r int) *TreeNode {
		// 當左右邊界不合法 (l > r)，表示子樹不存在
		if l > r {
			return nil
		}
		// 取出目前的根節點值，並讓 preIndex 向後移動一格
		rootVal := preorder[preIndex]
		preIndex++
		// 以 rootVal 建立當前樹的根節點
		root := &TreeNode{Val: rootVal}

		// 找出根節點在 inorder 中的位置
		inIndex := inMap[rootVal]

		// 以該位置將 inorder 分割成左右子樹，並遞歸建構
		root.Left = helper(l, inIndex-1)
		root.Right = helper(inIndex+1, r)

		return root
	}

	// 從 inorder 的最左 (0) 到最右 (len(inorder)-1) 範圍開始建構整棵樹
	return helper(0, len(inorder)-1)
}
