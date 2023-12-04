import unittest


def is_palindrome(s):
    front = 0
    rear = len(s) - 1

    while front <= rear:
        if s[front] != s[rear]:
            return False
        front += 1
        rear -= 1

    return True


class ValidPalindromeTest(unittest.TestCase):
    def test_a(self):
        self.assertTrue(is_palindrome('a'))

    def test_kayak(self):
        self.assertTrue(is_palindrome('kayak'))

    def test_hello(self):
        self.assertFalse(is_palindrome('hello'))

    def test_RACEACAR(self):
        self.assertFalse(is_palindrome('RACEACAR'))

    def test_ABCDABCD(self):
        self.assertFalse(is_palindrome('ABCDABCD'))
