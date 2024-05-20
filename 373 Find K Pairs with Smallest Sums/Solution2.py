import unittest
from typing import List
import heapq


class Solution:
    """
    You are given two integer arrays nums1 and nums2 sorted in non-decreasing order and an integer k.

    Define a pair (u, v) which consists of one element from the first array and one element from the second array.

    Return the k pairs (u1, v1), (u2, v2), ..., (uk, vk) with the smallest sums.



    Example 1:

    Input: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
    Output: [[1,2],[1,4],[1,6]]
    Explanation: The first 3 pairs are returned from the sequence: [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
    Example 2:

    Input: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
    Output: [[1,1],[1,1]]
    Explanation: The first 2 pairs are returned from the sequence: [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]


    Constraints:

    1 <= nums1.length, nums2.length <= 105
    -109 <= nums1[i], nums2[i] <= 109
    nums1 and nums2 both are sorted in non-decreasing order.
    1 <= k <= 104
    k <= nums1.length * nums2.length
    """

    def kSmallestPairs(self, nums1: List[int], nums2: List[int], k: int) -> List[List[int]]:
        k_smallest = []
        n1, n2 = len(nums1), len(nums2)
        heap = [(nums1[0] + nums2[0], 0, 0)]
        for _ in range(k):
            print(f"heap {heap}")
            sum, i, j = heapq.heappop(heap)
            # print(heap, i, j)
            print(f"get ({sum}, {i}, {j})")
            k_smallest.append([nums1[i], nums2[j]])

            if j == 0 and i + 1 < n1:
                print(f"A pushing ({nums1[i + 1] + nums2[j]}, {i+1}, {j})")
                heapq.heappush(heap, (nums1[i + 1] + nums2[j], i + 1, j))
            if j + 1 < n2:
                print(f"B pushing ({nums1[i] + nums2[j + 1]}, {i}, {j+1})")
                heapq.heappush(heap, (nums1[i] + nums2[j + 1], i, j + 1))
            print(f"heap {heap}")
            print("========")
        return k_smallest


class TestKPairsWithSmallestSums(unittest.TestCase):
    def setUp(self):
        self.solution = Solution()

    def test_example_1(self):
        nums1 = [1, 7, 11]
        nums2 = [2, 4, 6]
        k = 3
        self.assertListEqual([[1, 2], [1, 4], [1, 6]], self.solution.kSmallestPairs(nums1, nums2, k))

    def test_example_2(self):
        nums1 = [1, 1, 2]
        nums2 = [1, 2, 3]
        k = 2
        self.assertListEqual([[1, 1], [1, 1]], self.solution.kSmallestPairs(nums1, nums2, k))

    def test_example_3(self):
        nums1 = [1, 1, 2, 5]
        nums2 = [2, 3, 4]
        k = 3
        self.assertListEqual([[1, 2], [1, 2], [1, 3]], self.solution.kSmallestPairs(nums1, nums2, k))

    def test_example_4(self):
        nums1 = [2, 3, 4]
        nums2 = [1, 1, 2, 5]
        k = 5
        self.assertListEqual([[2, 1], [2, 1], [2, 2], [3, 1], [3, 1]], self.solution.kSmallestPairs(nums1, nums2, k))