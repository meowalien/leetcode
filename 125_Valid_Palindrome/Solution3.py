import unittest


class Solution:
    """
    A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

    Given a string s, return true if it is a palindrome, or false otherwise.



    Example 1:

    Input: s = "A man, a plan, a canal: Panama"
    Output: true
    Explanation: "amanaplanacanalpanama" is a palindrome.
    Example 2:

    Input: s = "race a car"
    Output: false
    Explanation: "raceacar" is not a palindrome.
    Example 3:

    Input: s = " "
    Output: true
    Explanation: s is an empty string "" after removing non-alphanumeric characters.
    Since an empty string reads the same forward and backward, it is a palindrome.


    Constraints:

    1 <= s.length <= 2 * 105
    s consists only of printable ASCII characters.
    """
    def isPalindrome(self, s: str) -> bool:
        s = ''.join(map(lambda x: x.lower() if x.isalnum() else '', s))
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
