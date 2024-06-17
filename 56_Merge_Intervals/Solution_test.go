package _6_Merge_Intervals

import "testing"

/*
Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.

Example 1:

Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
Example 2:

Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.

Constraints:

1 <= intervals.length <= 104
intervals[i].length == 2
0 <= starti <= endi <= 104
*/
func merge(intervals [][]int) [][]int {
	panic("implement me")
}

func testMerge(t *testing.T, result [][]int, correct [][]int) {
	if len(result) != len(correct) {
		t.Error("merge error")
	}
	for i := 0; i < len(result); i++ {
		if result[i][0] != correct[i][0] || result[i][1] != correct[i][1] {
			t.Error("merge error")
		}
	}
}

func TestMerge1(t *testing.T) {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	correct := [][]int{{1, 6}, {8, 10}, {15, 18}}

	result := merge(intervals)
	testMerge(t, result, correct)
}

func TestMerge2(t *testing.T) {
	intervals := [][]int{{1, 4}, {4, 5}}
	correct := [][]int{{1, 5}}

	result := merge(intervals)
	testMerge(t, result, correct)
}
