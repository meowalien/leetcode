package Insert_Interval

import (
	"testing"
)

/*
You are given an array of non-overlapping intervals intervals where intervals[i] = [starti, endi] represent the start and the end of the ith interval and intervals is sorted in ascending order by starti. You are also given an interval newInterval = [start, end] that represents the start and end of another interval.

Insert newInterval into intervals such that intervals is still sorted in ascending order by starti and intervals still does not have any overlapping intervals (merge overlapping intervals if necessary).

Return intervals after the insertion.

Note that you don't need to modify intervals in-place. You can make a new array and return it.

Example 1:

Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
Output: [[1,5],[6,9]]
Example 2:

Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
Output: [[1,2],[3,10],[12,16]]
Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].

Constraints:

0 <= intervals.length <= 104
intervals[i].length == 2
0 <= starti <= endi <= 105
intervals is sorted by starti in ascending order.
newInterval.length == 2
0 <= start <= end <= 105
*/
func insert1(intervals [][]int, newInterval []int) [][]int {
	var res = make([][]int, len(intervals)+1)
	i := 0

	// skip intervals end before the new one
	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		i++
	}
	cnt := i

	copy(res, intervals[:i])

	start, end := newInterval[0], newInterval[1]
	for i < len(intervals) && intervals[i][0] <= newInterval[1] {
		start = min(start, intervals[i][0])
		end = max(end, intervals[i][1])
		i++
	}

	res[cnt] = []int{start, end}

	copy(res[cnt+1:], intervals[i:])
	cnt += len(intervals) - i

	return res[:cnt+1]
}

func TestInsert_1(t *testing.T) {
	intervals := [][]int{{1, 3}, {6, 9}}
	newInterval := []int{2, 5}
	correct := [][]int{{1, 5}, {6, 9}}

	result := insert1(intervals, newInterval)
	if len(result) != len(correct) {
		t.Error("insert error")
	}
	for i := 0; i < len(result); i++ {
		if result[i][0] != correct[i][0] || result[i][1] != correct[i][1] {
			t.Error("insert error")
		}
	}
}

func TestInsert_11(t *testing.T) {
	intervals := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	newInterval := []int{4, 8}
	correct := [][]int{{1, 2}, {3, 10}, {12, 16}}

	result := insert1(intervals, newInterval)
	if len(result) != len(correct) {
		t.Error("insert error")
	}
	for i := 0; i < len(result); i++ {
		if result[i][0] != correct[i][0] || result[i][1] != correct[i][1] {
			t.Error("insert error")
		}
	}
}
