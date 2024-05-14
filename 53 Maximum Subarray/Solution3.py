import unittest
from typing import List


class Solution:
    """
    Given an integer array nums, find the
    subarray
     with the largest sum, and return its sum.



    Example 1:

    Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
    Output: 6
    Explanation: The subarray [4,-1,2,1] has the largest sum 6.
    Example 2:

    Input: nums = [1]
    Output: 1
    Explanation: The subarray [1] has the largest sum 1.
    Example 3:

    Input: nums = [5,4,-1,7,8]
    Output: 23
    Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.


    Constraints:

    1 <= nums.length <= 105
    -104 <= nums[i] <= 104


    Follow up: If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
    """

    def maxSubArray(self, nums: List[int]) -> int:
        if len(nums) == 1:
            return nums[0]

        for i in range(1, len(nums) * 2 - 1):
            if i < len(nums):
                if nums[i - 1] > 0:
                    nums[i] += nums[i - 1]
            else:
                i = i - len(nums) + 1
                nums[0] = max(nums[0], nums[i])
        return nums[0]


class TestMaxSubArray(unittest.TestCase):
    def setUp(self):
        self.solution = Solution()

    def test_max_sub_array1(self):
        # Test the example from the description
        self.assertEqual(self.solution.maxSubArray([-1, -2]), -1)

    def test_max_sub_array2(self):
        # Test the example from the description
        self.assertEqual(self.solution.maxSubArray([1, -1, -2]), 1)

    def test_max_sub_array(self):
        # Test the example from the description
        self.assertEqual(self.solution.maxSubArray([-2, 1, -3, 4, -1, 2, 1, -5, 4]), 6)

    def test_single_element(self):
        # Test an array with a single element
        self.assertEqual(self.solution.maxSubArray([1]), 1)

    def test_all_negative(self):
        # Test an array where all elements are negative
        self.assertEqual(self.solution.maxSubArray([-1, -2, -3, -4, -5]), -1)
