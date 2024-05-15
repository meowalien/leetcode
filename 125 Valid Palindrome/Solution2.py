import unittest


class Solution:
    def isPalindrome(self, s: str) -> bool:
        s = ''.join(map(lambda x: x.lower() if x.isalpha() or x.isdigit() else '', s))
        n = len(s)
        return s[:(n - 1) // 2 + 1] == s[n // 2:][::-1]


class TestIsPalindrome(unittest.TestCase):
    def setUp(self):
        self.solution = Solution()

    def test_isPalindrome1(self):
        result = self.solution.isPalindrome("A man, a plan, a canal: Panama")
        self.assertEqual(result, True)

    def test_isPalindrome2(self):
        result = self.solution.isPalindrome("race a car")
        self.assertEqual(result, False)

    def test_isPalindrome3(self):
        result = self.solution.isPalindrome("12321")
        self.assertEqual(result, True)

    def test_isPalindrome4(self):
        result = self.solution.isPalindrome("abcba")
        self.assertEqual(result, True)
