import unittest
from typing import List


class Solution:
    """
    According to Wikipedia's article: "The Game of Life, also known simply as Life, is a cellular automaton devised by the British mathematician John Horton Conway in 1970."

    The board is made up of an m x n grid of cells, where each cell has an initial state: live (represented by a 1) or dead (represented by a 0). Each cell interacts with its eight neighbors (horizontal, vertical, diagonal) using the following four rules (taken from the above Wikipedia article):

    Any live cell with fewer than two live neighbors dies as if caused by under-population.
    Any live cell with two or three live neighbors lives on to the next generation.
    Any live cell with more than three live neighbors dies, as if by over-population.
    Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
    The next state is created by applying the above rules simultaneously to every cell in the current state, where births and deaths occur simultaneously. Given the current state of the m x n grid board, return the next state.



    Example 1:


    Input: board = [[0,1,0],[0,0,1],[1,1,1],[0,0,0]]
    Output: [[0,0,0],[1,0,1],[0,1,1],[0,1,0]]
    Example 2:


    Input: board = [[1,1],[1,0]]
    Output: [[1,1],[1,1]]


    Constraints:

    m == board.length
    n == board[i].length
    1 <= m, n <= 25
    board[i][j] is 0 or 1.


    Follow up:

    Could you solve it in-place? Remember that the board needs to be updated simultaneously: You cannot update some cells first and then use their updated values to update other cells.
    In this question, we represent the board using a 2D array. In principle, the board is infinite, which would cause problems when the active area encroaches upon the border of the array (i.e., live cells reach the border). How would you address these problems?
    """

    def gameOfLife(self, board: List[List[int]]) -> None:
        """
        Do not return anything, modify board in-place instead.
        """

        def update_cell(j, i):
            life_count = 0
            for x in range(max(0, i - 1), min(len(board[0]), i + 2)):
                for y in range(max(0, j - 1), min(len(board), j + 2)):
                    if x == i and y == j:
                        continue
                    life_count += board[y][x] & 1
            match life_count:
                case 3:
                    life_count = 1
                case 2:
                    life_count = board[j][i]
                case _:
                    life_count = 0
            # print(f"visit:({j}, {i}) {board[j][i]} -> {life_count}")
            next_i = i + 1 if i + 1 < len(board[0]) else 0
            if next_i != 0 and next_i > i:
                next_j = j
            else:
                next_j = j + 1 if j + 1 < len(board) else -1

            if next_j != -1:
                update_cell(next_j, next_i)
            board[j][i] = life_count

        update_cell(0, 0)


class TestGameOfLife(unittest.TestCase):
    def setUp(self):
        self.solution = Solution()

    def test_gameOfLife(self):
        board = [[0, 1, 0], [0, 0, 1], [1, 1, 1], [0, 0, 0]]
        self.solution.gameOfLife(board)
        self.assertEqual(board, [[0, 0, 0], [1, 0, 1], [0, 1, 1], [0, 1, 0]])

    def test_gameOfLife1(self):
        board = [[1, 1], [1, 0]]
        self.solution.gameOfLife(board)
        self.assertEqual(board, [[1, 1], [1, 1]])

    def test_gameOfLife2(self):
        board = [[1, 1, 1, 1], [1, 1, 1, 1], [1, 1, 1, 1], [1, 1, 1, 1]]
        self.solution.gameOfLife(board)
        self.assertEqual(board, [[1, 0, 0, 1], [0, 0, 0, 0], [0, 0, 0, 0], [1, 0, 0, 1]])
