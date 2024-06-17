import unittest
from typing import List


class Solution:
    """
    Given an array of points where points[i] = [xi, yi] represents a point on the X-Y plane, return the maximum number of points that lie on the same straight line.



    Example 1:


    Input: points = [[1,1],[2,2],[3,3]]
    Output: 3
    Example 2:


    Input: points = [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
    Output: 4


    Constraints:

    1 <= points.length <= 300
    points[i].length == 2
    -104 <= xi, yi <= 104
    All the points are unique.
    """

    def maxPoints(self, points: List[List[int]]) -> int:
        if len(points) == 1:
            return 1

        max_points = 0
        for i in range(len(points)):
            slope_count = {}
            for j in range(i + 1, len(points)):
                slope = self.calculate_slope(points[i], points[j])
                slope_count[slope] = slope_count.get(slope, 1) + 1
            max_points = max(max_points, max(slope_count.values(), default=0))

        return max_points

    def calculate_slope(self, current_point: List[int], param: List[int]) -> float:
        x1, y1 = current_point
        x2, y2 = param
        if x1 == x2:
            return float('inf')
        if y1 == y2:
            return 0
        a = y2 - y1
        b = x2 - x1

        gcd = self.gcd(a, b)
        return (a // gcd) / (b // gcd)

    def gcd(self, a, b):
        while b != 0:
            a, b = b, a % b
        return abs(a)


class TestMaxPoints(unittest.TestCase):
    def setUp(self):
        self.solution = Solution()

    def test_example_1(self):
        points = [[1, 1], [2, 2], [3, 3]]
        self.assertEqual(3, self.solution.maxPoints(points))

    def test_example_2(self):
        points = [[1, 1], [3, 2], [5, 3], [4, 1], [2, 3], [1, 4]]
        self.assertEqual(4, self.solution.maxPoints(points))

    def test_example_3(self):
        points = [[0, 0]]
        self.assertEqual(1, self.solution.maxPoints(points))

    def test_example_4(self):
        points = [[4, 5], [4, -1], [4, 0]]
        self.assertEqual(3, self.solution.maxPoints(points))

    def test_example_5(self):
        points = [[0, 0], [1, -1], [1, 1]]
        self.assertEqual(2, self.solution.maxPoints(points))

    def test_example_6(self):
        points = [[1, 1], [2, 2], [3, 3]]
        self.assertEqual(3, self.solution.maxPoints(points))


class TestSlope(unittest.TestCase):
    def setUp(self):
        self.solution = Solution()

    def test_slope_1(self):
        a = [1, 1]
        b = [2, 2]
        ans = self.solution.calculate_slope(a, b)
        self.assertEqual(ans, (1, 1))

    def test_slope_2(self):
        a = [12, 6]
        b = [2, 1]
        ans = self.solution.calculate_slope(a, b)
        self.assertEqual(ans, (2, 1))

    def test_slope_2_1(self):
        a = [2, 1]
        b = [0, 0]
        ans = self.solution.calculate_slope(a, b)
        self.assertEqual(ans, (2, 1))

    def test_slope_3(self):
        a = [1, 1]
        b = [1, 2]
        ans = self.solution.calculate_slope(a, b)
        self.assertEqual(ans, (0, 1))

    def test_slope_4(self):
        a = [1, 1]
        b = [2, 1]
        ans = self.solution.calculate_slope(a, b)
        self.assertEqual(ans, (1, 0))

    def test_slope_5(self):
        a = [17, 9]
        b = [5, 5]
        ans = self.solution.calculate_slope(a, b)
        self.assertEqual(ans, (3, 1))

    def test_slope_6(self):
        a = [5, 5]
        b = [17, 9]
        ans = self.solution.calculate_slope(a, b)
        self.assertEqual(ans, (3, 1))

    def test_slope_7(self):
        a = [0, 1]
        b = [0, 9]
        ans = self.solution.calculate_slope(a, b)
        self.assertEqual(ans, (3, 1))

    def test_slope_8(self):
        a = [0, 1]
        b = [0, 5]
        ans = self.solution.calculate_slope(a, b)
        self.assertEqual(ans, (3, 1))
