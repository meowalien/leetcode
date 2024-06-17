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
        board_copy = [row[:] for row in board]

        m = len(board)
        n = len(board[0])
        b2 = [[0 for i in range(n)] for _ in range(m)]
        for row in range(len(board)):
            for block in range(len(board[row])):
                c = 0
                if board_copy[row][block] == 1:
                    c += 1
                #block right
                if block+1 < len(board[row]):
                    if board_copy[row][block+1] == 1:
                        c += 1
                #block left
                if block>0:
                    if board_copy[row][block-1] == 1:
                        c += 1
                #block below
                if row+1 < len(board):
                    if board_copy[row+1][block] == 1:
                        c +=1
                #block down/right
                if row+1 < len(board) and block+1 < len(board[row]):
                    if board_copy[row+1][block+1] == 1:
                        c += 1
                #block down/left
                if row+1 < len(board) and block > 0:
                    if board_copy[row+1][block-1] == 1:
                        c += 1
                #block above
                if row-1 >= 0:
                    if board_copy[row-1][block] == 1:
                        c += 1
                #block up/right
                if row-1 >= 0 and block+1 < len(board[row]):
                    if board_copy[row-1][block+1] == 1:
                        c += 1
                #block up/left
                if row-1 >= 0 and block > 0:
                    if board_copy[row-1][block-1] == 1:
                        c += 1
                #rule 1
                if board_copy[row][block] == 1 and c <= 2:
                    new_block = 0
                #rule 2
                elif board_copy[row][block] == 1 and 3 <= c <=4:
                    new_block = 1
                #rule 3
                elif board_copy[row][block] == 1 and c > 4:
                    new_block = 0
                #rule 4
                elif board_copy[row][block] == 0 and c == 3:
                    new_block = 1
                #dead cell
                elif board_copy[row][block] == 0:
                    new_block = 0
                board[row][block] = new_block


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
