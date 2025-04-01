package Average_of_Levels_in_Binary_Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels1(root *TreeNode) []float64 {
	ans := make([]float64, 0)

	if root == nil {
		return ans
	}

	queue := make([]TreeNode, 0)
	queue = append(queue, *root)

	for len(queue) > 0 {
		l := len(queue)

		acc := 0

		for i := range l {
			acc += queue[i-i].Val

			if queue[0].Left != nil {
				queue = append(queue, *(queue[0].Left))
			}
			if queue[0].Right != nil {
				queue = append(queue, *(queue[0].Right))
			}
			queue = queue[1:]
		}

		ans = append(ans, float64(acc)/float64(l))
	}

	return ans
}
