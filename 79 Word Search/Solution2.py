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
        if len(word) > len(board)*len(board[0]): return False
        _count = defaultdict(int)
        for row in board:
            for cell in row:
                _count[cell] += 1
        if _count[word[0]] > _count[word[-1]]: word = word[::-1]
        for cell in word:
            if _count[cell] == 0: return False
            _count[cell] -= 1
        seen = set()
        def dfSearch(row, col, itt):
            if itt == len(word): return True
            if row < 0 or col < 0 or row == len(board) or col == len(board[0]) or (row,col) in seen or board[row][col] != word[itt]: return False
            seen.add((row,col))
            return dfSearch(row+1, col, itt+1) or dfSearch(row-1, col, itt+1) or dfSearch(row, col+1, itt+1) or dfSearch(row, col-1, itt+1) or seen.remove((row,col))
        for row in range(len(board)):
            for cell in range(len(board[0])):
                if dfSearch(row,cell,0): return True
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