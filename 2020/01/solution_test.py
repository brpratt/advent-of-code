import unittest

from .solution import find_sum_pair, find_sum_thrice, solve_part_1, solve_part_2


class TestFindSumPair(unittest.TestCase):
    def test_returns_pair(self):
        nums = [1721, 979, 366, 299, 675, 1456]
        expectedPair = (1721, 299)
        self.assertEqual(find_sum_pair(nums, 2020), expectedPair)

    def test_raises_exception_on_no_pair(self):
        nums = [1721, 979, 366, 299, 675, 1456]
        self.assertRaises(Exception, lambda _: find_sum_pair(nums, 3))


class TestFindSumThrice(unittest.TestCase):
    def test_returns_thrice(self):
        nums = [1721, 979, 366, 299, 675, 1456]
        expectedThrice = (979, 366, 675)
        self.assertEqual(find_sum_thrice(nums, 2020), expectedThrice)


class TestSolvePart1(unittest.TestCase):
    def test_returns_answer(self):
        nums = [1721, 979, 366, 299, 675, 1456]
        self.assertEqual(solve_part_1(nums), 514579)


class TestSolvePart2(unittest.TestCase):
    def test_returns_answer(self):
        nums = [1721, 979, 366, 299, 675, 1456]
        self.assertEqual(solve_part_2(nums), 241861950)
