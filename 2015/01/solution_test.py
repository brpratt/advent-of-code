import unittest

from .solution import solve_part_1, solve_part_2


class TestDay01(unittest.TestCase):
    def test_solve_part_1(self) -> None:
        tests = [
            ("(())", 0),
            ("()()", 0),
            ("(((", 3),
            ("(()(()(", 3),
            ("))(((((", 3),
            ("())", -1),
            ("))(", -1),
            (")))", -3),
            (")())())", -3),
        ]

        for line, expected in tests:
            self.assertEqual(solve_part_1(line), expected)

    def test_solve_part_2(self) -> None:
        tests = [(")", 1), ("()())", 5)]

        for line, expected in tests:
            self.assertEqual(solve_part_2(line), expected)

    def test_answers(self) -> None:
        with open("./2015/01/input.txt", "r") as f:
            line = f.readline()
            self.assertEqual(solve_part_1(line), 74)
            self.assertEqual(solve_part_2(line), 1795)


if __name__ == "__main__":
    unittest.main()
