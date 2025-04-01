package Average_of_Levels_in_Binary_Tree

/*
Given the root of a binary tree, return the average value of the nodes on each level in the form of an array. Answers within 10-5 of the actual answer will be accepted.

Example 1:

Input: root = [3,9,20,null,null,15,7]
Output: [3.00000,14.50000,11.00000]
Explanation: The average value of nodes on level 0 is 3, on level 1 is 14.5, and on level 2 is 11.
Hence return [3, 14.5, 11].
Example 2:

Input: root = [3,9,20,15,7]
Output: [3.00000,14.50000,11.00000]

Constraints:

The number of nodes in the tree is in the range [1, 104].
-231 <= Node.val <= 231 - 1
*/
func averageOfLevels(root *TreeNode) []float64 {
	queue := []*TreeNode{root}
	res := []float64{}
	//visited := map[*TreeNode]struct{}{}
	for len(queue) > 0 {
		sizeForThisLevel := len(queue)
		sum := 0
		for i := 0; i < sizeForThisLevel; i++ {
			node := queue[0]
			queue = queue[1:]
			//if _, ok := visited[node]; ok {
			//	continue
			//}
			//visited[node] = struct{}{}
			sum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, float64(sum)/float64(sizeForThisLevel))
	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
