import unittest
from collections import defaultdict
from typing import List


class Solution:
    """
    Given an m x n grid of characters board and a string word, return true if word exists in the grid.

    The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once.



    Example 1:


    Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
    Output: true
    Example 2:


    Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
    Output: true
    Example 3:


    Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
    Output: false


    Constraints:

    m == board.length
    n = board[i].length
    1 <= m, n <= 6
    1 <= word.length <= 15
    board and word consists of only lowercase and uppercase English letters.


    Follow up: Could you use search pruning to make your solution faster with a larger board?
    """

    def exist(self, board: List[List[str]], word: str) -> bool:
        def backtrack(i, j, k):
            if k == len(word):
                return True
            if i < 0 or i >= len(board) or j < 0 or j >= len(board[0]) or board[i][j] != word[k]:
                return False

            temp = board[i][j]
            board[i][j] = ''

            if backtrack(i + 1, j, k + 1) or backtrack(i - 1, j, k + 1) or backtrack(i, j + 1, k + 1) or backtrack(i,
                                                                                                                   j - 1,
                                                                                                                   k + 1):
                return True

            board[i][j] = temp
            return False

        for i in range(len(board)):
            for j in range(len(board[0])):
                if backtrack(i, j, 0):
                    return True
        return False

    def print_board(self, board):
        for row in board:
            print(row)
        print()


class TestWordSearch(unittest.TestCase):
    def setUp(self):
        self.solution = Solution()

    def test_example_1(self):
        board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]
        word = "ABCCED"
        self.assertTrue(self.solution.exist(board, word))

    def test_example_2(self):
        board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]
        word = "SEE"
        self.assertTrue(self.solution.exist(board, word))

    def test_example_3(self):
        board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]
        word = "ABCB"
        self.assertFalse(self.solution.exist(board, word))