package _20_Triangle

/*
Given a triangle array, return the minimum path sum from top to bottom.

For each step, you may move to an adjacent number of the row below. More formally, if you are on index i on the current row, you may move to either index i or index i + 1 on the next row.

Example 1:

Input: triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
Output: 11
Explanation: The triangle looks like:

	  2
	 3 4
	6 5 7

4 1 8 3
The minimum path sum from top to bottom is 2 + 3 + 5 + 1 = 11 (underlined above).
Example 2:

Input: triangle = [[-10]]
Output: -10

Constraints:

1 <= triangle.length <= 200
triangle[0].length == 1
triangle[i].length == triangle[i - 1].length + 1
-104 <= triangle[i][j] <= 104

Follow up: Could you do this using only O(n) extra space, where n is the total number of rows in the triangle?
*/
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 1 {
		return triangle[0][0]
	}

	return getTheMinSum(triangle, 0, 0, 0)
}

func getTheMinSum(triangle [][]int, rowIndex, colIndex, sum int) int {
	current := triangle[rowIndex][colIndex] + sum
	//println("current: ", current)
	if rowIndex == len(triangle)-1 {
		return current
	}

	left := getTheMinSum(triangle, rowIndex+1, colIndex, current)
	right := getTheMinSum(triangle, rowIndex+1, colIndex+1, current)
	return min(left, right)
}

func printTriangle(printTriangle [][]int) {
	for _, row := range printTriangle {
		for _, value := range row {
			print(value, " ")
		}
		println()
	}
}