package Minimum_Size_Subarray_Sum

import (
	"math"
	"testing"
)

/*
Given an array of positive integers nums and a positive integer target, return the minimal length of a
subarray

	whose sum is greater than or equal to target. If there is no such subarray, return 0 instead.

Example 1:

Input: target = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: The subarray [4,3] has the minimal length under the problem constraint.
Example 2:

Input: target = 4, nums = [1,4,4]
Output: 1
Example 3:

Input: target = 11, nums = [1,1,1,1,1,1,1,1]
Output: 0

Constraints:

1 <= target <= 109
1 <= nums.length <= 105
1 <= nums[i] <= 104

Follow up: If you have figured out the O(n) solution, try coding another solution of which the time complexity is O(n log(n)).
*/

func minSubArrayLen(target int, nums []int) int {
	left, right, sum, minLen := 0, 0, 0, math.MaxInt32

	for right < len(nums) {
		//fmt.Println("left:", left, "right:", right, "sum:", sum, "minLen:", minLen)
		sum += nums[right]
		for sum >= target {
			minLen = min(minLen, right-left+1)
			sum -= nums[left]
			left++
		}
		right++
	}

	if minLen == math.MaxInt32 {
		return 0
	}
	return minLen
}

func TestSolution1(t *testing.T) {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	result := minSubArrayLen(target, nums)
	if result != 2 {
		t.Error("minSubArrayLen error")
	}
}

func TestSolution2(t *testing.T) {
	target := 4
	nums := []int{1, 4, 4}
	result := minSubArrayLen(target, nums)
	if result != 1 {
		t.Error("minSubArrayLen error")
	}
}

func TestSolution3(t *testing.T) {
	target := 11
	nums := []int{1, 1, 1, 1, 1, 1, 1, 1}
	result := minSubArrayLen(target, nums)
	if result != 0 {
		t.Error("minSubArrayLen error")
	}
}
