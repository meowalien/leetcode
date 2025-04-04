package Find_First_and_Last_Position_of_Element_in_Sorted_Array

/*
Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.

If target is not found in the array, return [-1, -1].

You must write an algorithm with O(log n) runtime complexity.

Example 1:

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
Example 2:

Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
Example 3:

Input: nums = [], target = 0
Output: [-1,-1]

Constraints:

0 <= nums.length <= 105
-109 <= nums[i] <= 109
nums is a non-decreasing array.
-109 <= target <= 109
*/
func searchRange(nums []int, target int) []int {
	first := findFirst(nums, target)
	last := findLast(nums, target)
	return []int{first, last}
}

func findFirst(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := -1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			result = mid
			right = mid - 1 // keep searching in the left half
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

func findLast(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := -1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			result = mid
			left = mid + 1 // keep searching in the right half
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}
