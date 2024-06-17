import unittest
from typing import List


class Solution:
    """
    You are given an integer array prices where prices[i] is the price of a given stock on the ith day, and an integer k.

    Find the maximum profit you can achieve. You may complete at most k transactions: i.e. you may buy at most k times and sell at most k times.

    Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).



    Example 1:

    Input: k = 2, prices = [2,4,1]
    Output: 2
    Explanation: Buy on day 1 (price = 2) and sell on day 2 (price = 4), profit = 4-2 = 2.
    Example 2:

    Input: k = 2, prices = [3,2,6,5,0,3]
    Output: 7
    Explanation: Buy on day 2 (price = 2) and sell on day 3 (price = 6), profit = 6-2 = 4. Then buy on day 5 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.


    Constraints:

    1 <= k <= 100
    1 <= prices.length <= 1000
    0 <= prices[i] <= 1000
    """

    def maxProfit(self, k: int, prices: List[int]) -> int:
        """
        the best strategy is to buy at the lowest price and sell at the highest price
        but we can only buy and sell at most k times

        so we need to do the best transactions first, then the second best, and so on
        the best transaction is the one with the highest profit, in other words,
        the one with the highest difference between the buy and sell prices
        """
        print("prices: ", prices)
        # if there are no prices, we can't make any profit
        if not prices:
            return 0

        # find all the rising sequences in the prices list and order them by profit
        rising_sequences = self.find_rising_sequences(prices)

        if k < len(rising_sequences):
            self.merge_the_sequences(k, prices, rising_sequences)
        print("result after merge: ", rising_sequences)
        # order the sequences by profit
        rising_sequences.sort(key=lambda x: x[2], reverse=True)

        print("rising_sequences: ", rising_sequences)
        profit = 0

        doable_transactions = min(k, len(rising_sequences))
        # do the transactions
        for i in range(doable_transactions):
            # print("i: ", i)
            profit += rising_sequences[i][2]
        return profit

    def find_rising_sequences(self, prices: List[int]) -> List[List[int]]:
        result = []
        # the first element in the sequence is the buy price's index
        # the second element in the sequence is the sell price's index
        # the third element in the sequence is the profit
        current_sequences = [0, 0, 0]

        for i in range(1, len(prices)):
            # print("current_sequences: ",current_sequences)
            highest_price_in_sequence = prices[current_sequences[1]]
            current_price = prices[i]
            if current_price >= highest_price_in_sequence:
                # print("adding to the sequence")
                current_sequences[1] = i
                # print("current_sequences: ", current_sequences)
                if i < len(prices) - 1:
                    continue

            # calculate the profit
            current_sequences[2] = prices[current_sequences[1]] - prices[current_sequences[0]]
            # if the profit is positive, add it to the result
            if current_sequences[2] > 0:
                result.append(current_sequences)

            # start a new sequence
            current_sequences = [i, i, 0]

        print("result: ", result)

        return result

    def merge_the_sequences(self, k: int, prices: List[int], result: List[List[int]]):
        if len(result) <= k:
            return
        for x in range(1, len(result) - 1):
            # print("result: ", result)
            # print("x: ", x)
            the_lowest_price = prices[result[x - 1][0]]
            the_middle_price = prices[result[x][0]]
            the_highest_price = prices[result[x][1]]
            # print("the_lowest_price: ", the_lowest_price)
            # print("the_middle_price: ", the_middle_price)
            # print("the_highest_price: ", the_highest_price)
            # print("=====================================")
            if the_lowest_price <= the_middle_price <= the_highest_price:
                #         merge the two sequences
                result[x - 1][1] = result[x][1]
                result[x - 1][2] = prices[result[x][1]] - prices[result[x - 1][0]]
                result.pop(x)
                self.merge_the_sequences(k, prices, result)

                return
            # print("result after merge: ", result)
            # print("result after merge: ", result)


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

    def test_1(self):
        # Mixed profitable and non-profitable transactions
        # Buy at 1 sell at 5, buy at 3 sell at 6
        self.assertEqual(self.solution.maxProfit(2, [1, 2, 4, 2, 5, 7, 2, 4, 9, 0]), 13)

    def test_3(self):
        # Mixed profitable and non-profitable transactions
        # Buy at 1 sell at 5, buy at 3 sell at 6
        self.assertEqual(self.solution.maxProfit(4, [1, 2, 4, 2, 5, 7, 2, 4, 9, 0]), 15)

    def test_2(self):
        # Mixed profitable and non-profitable transactions
        # Buy at 1 sell at 5, buy at 3 sell at 6
        self.assertEqual(self.solution.maxProfit(2, [1, 2, 4, 2, 5, 7, 2, 4, 9, 0, 9]), 17)

    def test_4(self):
        # Mixed profitable and non-profitable transactions
        # Buy at 1 sell at 5, buy at 3 sell at 6
        self.assertEqual(self.solution.maxProfit(2, [8, 6, 4, 3, 3, 2, 3, 5, 8, 3, 8, 2, 6]), 11)


if __name__ == "__main__":
    unittest.main()
