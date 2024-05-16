# Definition for a binary tree node.
import unittest
from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    """
    Given the root of a binary search tree, and an integer k, return the kth smallest value (1-indexed) of all the values of the nodes in the tree.



    Example 1:


    Input: root = [3,1,4,null,2], k = 1
    Output: 1
    Example 2:


    Input: root = [5,3,6,2,4,null,null,1], k = 3
    Output: 3


    Constraints:

    The number of nodes in the tree is n.
    1 <= k <= n <= 104
    0 <= Node.val <= 104


    Follow up: If the BST is modified often (i.e., we can do insert and delete operations) and you need to find the kth smallest frequently, how would you optimize?
    """

    def kthSmallest(self, root: Optional[TreeNode], k: int) -> int:
        if not root:
            return 0
        current = root

        while current:
            if current.left:
                pre = current.left
                while pre.right and pre.right != current:
                    pre = pre.right
                if not pre.right:
                    pre.right = current
                    current = current.left
                else:
                    k -= 1
                    if k == 0:
                        return current.val
                    pre.right = None
                    current = current.right
            else:
                k -= 1
                if k == 0:
                    return current.val
                current = current.right


class TestKthSmallest(unittest.TestCase):
    def setUp(self):
        self.solution = Solution()

    def make_tree(self, nodes):
        if not nodes:
            return None
        root = TreeNode(nodes[0])
        queue = [root]
        i = 1
        while queue and i < len(nodes):
            node = queue.pop(0)
            if nodes[i] is not None:
                node.left = TreeNode(nodes[i])
                queue.append(node.left)
            i += 1
            if i < len(nodes) and nodes[i] is not None:
                node.right = TreeNode(nodes[i])
                queue.append(node.right)
            i += 1
        return root

    def test_kthSmallest1(self):
        root = self.make_tree([3, 1, 4, None, 2])
        result = self.solution.kthSmallest(root, 1)
        self.assertEqual(result, 1)

    def test_kthSmallest2(self):
        root = self.make_tree([5, 3, 6, 2, 4, None, None, 1])
        result = self.solution.kthSmallest(root, 3)
        self.assertEqual(result, 3)

    def test_kthSmallest3(self):
        root = self.make_tree([2, 1])
        result = self.solution.kthSmallest(root, 2)
        self.assertEqual(result, 2)

    def test_kthSmallest4(self):
        root = self.make_tree([1, None, 2])
        result = self.solution.kthSmallest(root, 2)
        self.assertEqual(result, 2)

    def test_kthSmallest5(self):
        root = self.make_tree([50, 30, 70, 20, 40, 60, 80])
        result = self.solution.kthSmallest(root, 5)
        self.assertEqual(result, 60)
