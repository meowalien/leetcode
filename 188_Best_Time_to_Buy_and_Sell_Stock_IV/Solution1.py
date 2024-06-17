import unittest
from typing import List


class Solution:
    def maxProfit(self, k: int, prices: List[int]) -> int:
        print(f"k:{k}, prices:{prices}")
        if not prices:
            return 0
        n = len(prices)
        if k >= n // 2:
            return sum(max(prices[i + 1] - prices[i], 0) for i in range(n - 1))
        dp = [[0] * n for _ in range(k + 1)]
        for i in range(1, k + 1):
            # the cost of buying the stock at the first day
            max_diff = -prices[0]
            # print("max_diff: ", max_diff)
            for j in range(1, n):
                print("i: ", i)
                print("j: ", j)
                dp[i][j] = max(dp[i][j - 1], prices[j] + max_diff)
                print("dp[i][j - 1]: ",dp[i][j - 1])
                print("max_diff: ",max_diff)
                print("prices[j]: ",prices[j])
                max_diff = max(max_diff, dp[i - 1][j] - prices[j])
                print("max_diff update: ",max_diff)
                print("dp[i - 1][j]: ", dp[i - 1][j])
                print("prices[j]: ", prices[j])

                print( prices)
                for x in dp:
                    print(x)
                print("=====================")
                # print("max_diff: ", max_diff)
            # print("dp: ",dp)
        print("result:", dp[k][n - 1])
        return dp[k][n - 1]


class TestMaxProfit(unittest.TestCase):
    def setUp(self):
        self.solution = Solution()

    def test_no_transactions(self):
        # Test with zero transactions allowed
        self.assertEqual(self.solution.maxProfit(0, [3, 2, 6, 5, 0, 3]), 0)

    def test_no_prices(self):
        # Test with an empty prices list
        self.assertEqual(self.solution.maxProfit(1, []), 0)

    def test_multiple_transactions(self):
        # Test with multiple transactions allowed more than needed
        # Buy at 1 and sell at 2 and buy at 2 and sell at 3 and buy at 3 and sell at 4 and buy at 4 and sell at 5
        self.assertEqual(self.solution.maxProfit(3, [1, 2, 3, 4, 5]), 4)

    def test_limited_transactions(self):
        # Test with limited transactions where selection is crucial
        # Buy at 2 and sell at 6, buy at 0 and sell at 3
        self.assertEqual(self.solution.maxProfit(2, [3, 2, 6, 5, 0, 3]), 7)

    def test_k_greater_than_n_div_2(self):
        # Test the edge case where k is greater than half of the number of days
        # Buy at 10 sell at 22, buy at 5 sell at 75 and buy at 65 sell at 80
        self.assertEqual(self.solution.maxProfit(10, [10, 22, 5, 75, 65, 80]), 97)

    def test_profitable_and_nonprofitable_mixed(self):
        # Mixed profitable and non-profitable transactions
        # Buy at 1 sell at 5, buy at 3 sell at 6
        self.assertEqual(self.solution.maxProfit(2, [7, 1, 3, 5, 3, 6, 4]), 7)


if __name__ == "__main__":
    unittest.main()
