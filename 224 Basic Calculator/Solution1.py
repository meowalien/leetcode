import unittest


class Solution:
    """
    Given a string s representing a valid expression, implement a basic calculator to evaluate it, and return the result of the evaluation.

    Note: You are not allowed to use any built-in function which evaluates strings as mathematical expressions, such as eval().



    Example 1:

    Input: s = "1 + 1"
    Output: 2
    Example 2:

    Input: s = " 2-1 + 2 "
    Output: 3
    Example 3:

    Input: s = "(1+(4+5+2)-3)+(6+8)"
    Output: 23


    Constraints:

    1 <= s.length <= 3 * 105
    s consists of digits, '+', '-', '(', ')', and ' '.
    s represents a valid expression.
    '+' is not used as a unary operation (i.e., "+1" and "+(2 + 3)" is invalid).
    '-' could be used as a unary operation (i.e., "-1" and "-(2 + 3)" is valid).
    There will be no two consecutive operators in the input.
    Every number and running calculation will fit in a signed 32-bit integer.
    """

    def calculate(self, s: str) -> int:
        stack = []
        result = 0
        num = 0
        sign = 1
        for c in s:
            if c.isdigit():
                num = num * 10 + int(c)
            elif c == '+':
                result += sign * num
                num = 0
                sign = 1
            elif c == '-':
                result += sign * num
                num = 0
                sign = -1
            elif c == '(':
                stack.append(result)
                stack.append(sign)
                result = 0
                sign = 1
            elif c == ')':
                result += sign * num
                result *= stack.pop()
                result += stack.pop()
                num = 0
            print("c: ", c)
            print("stack: ", stack)
            print("num: ", num)
            print("result: ", result)
            print("sign: ", sign)
            print("=====================================")

        return result + sign * num


class TestCalculate(unittest.TestCase):
    def setUp(self):
        self.solution = Solution()

    def test_example_0(self):
        algorithm = "1 + 1"
        answer = eval(algorithm)
        self.assertEqual(self.solution.calculate(algorithm), answer)

    def test_example_1(self):
        algorithm = "1864 + 1"
        answer = eval(algorithm)
        self.assertEqual(self.solution.calculate(algorithm), answer)

    def test_example_2(self):
        algorithm = " 2-1 + 2 "
        answer = eval(algorithm)
        self.assertEqual(self.solution.calculate(algorithm), answer)

    def test_example_3(self):
        algorithm = "(1+(4+5+2-3)-3)+(6+8)"
        answer = eval(algorithm)
        self.assertEqual(self.solution.calculate(algorithm), answer)

    def test_example_4(self):
        algorithm = "(1-204+(4+599+2-33)-53)+(6-8)"
        answer = eval(algorithm)
        self.assertEqual(self.solution.calculate(algorithm), answer)
