import unittest
from heapq import heappush, heappop
from typing import List


class DoubleLinkListNode:
    def __init__(self, ind, pre=None, next=None):
        self.ind = ind
        self.pre = pre if pre else self
        self.next = next if next else self


def print_double_linked_list(head, newP):
    """
    Prints the indices of the nodes in the double linked list starting from the given head.

    :param head: The head node of the double linked list
    """
    current = head
    while current is not None:
        print("" if current.ind == -1 else newP[current.ind], end="" if current.ind == -1 or current.next.ind == -1 else ' -> ')
        current = current.next
        if current.ind == -1:
            break
    print("")


class Solution:
    def MinMaxList(self, arr: List[int]) -> List[int]:
        print("arr: ", arr)
        n = len(arr)
        if n == 0:
            return []
        sign = -1
        res = [9999]
        for num in arr:
            if num * sign > res[-1] * sign:
                res[-1] = num
            else:
                res.append(num)
                sign *= -1
        if len(res) & 1:
            res.pop()
        return res

    def maxProfit(self, k: int, prices: List[int]) -> int:
        newP = self.MinMaxList(prices)
        print("newP: ", newP)
        n = len(newP)
        m = n // 2
        print("m: ", m)
        print("k: ", k)
        res = 0
        for i in range(m):
            res += newP[i * 2 + 1] - newP[i * 2]
        if m <= k:
            return res

        head, tail = DoubleLinkListNode(-1), DoubleLinkListNode(-1)
        NodeList = [DoubleLinkListNode(0, head)]
        for i in range(1, n):
            NodeList.append(DoubleLinkListNode(i, NodeList[-1]))
            NodeList[i - 1].next = NodeList[i]
        NodeList[n - 1].next = tail
        head.next, tail.pre = NodeList[0], NodeList[n - 1]
        print("=====================================")
        print_double_linked_list(head, newP)
        heap = []
        for i in range(n - 1):
            if i & 1:
                heappush(heap, [newP[i] - newP[i + 1], i, i + 1, 0])
            else:
                heappush(heap, [newP[i + 1] - newP[i], i, i + 1, 1])
        while m > k:
            loss, i, j, t = heappop(heap)
            if NodeList[i] == None or NodeList[j] == None: continue
            m -= 1
            res -= loss
            print("loss: ", loss)
            nodei, nodej = NodeList[i], NodeList[j]
            print("deleting: ", newP[nodei.ind],  newP[nodej.ind])
            nodel, noder = nodei.pre, nodej.next
            l, r = nodel.ind, noder.ind
            valL, valR = newP[l], newP[r]
            noder.pre, nodel.next = nodel, noder
            print_double_linked_list(head, newP)
            NodeList[i], NodeList[j] = None, None
            if t == 0:
                heappush(heap, [valR - valL, l, r, 1])
            elif l != -1 and r != -1:
                heappush(heap, [valL - valR, l, r, 0])

        return res


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
