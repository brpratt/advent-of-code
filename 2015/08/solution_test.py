import unittest

from .solution import solve_part_1, solve_part_2


class TestDay08(unittest.TestCase):
    def test_solve_part_1(self) -> None:
        contents = """""
"abc"
"aaa\\"aaa"
"\\x27"
"""

        lines = contents.split("\n")
        self.assertEqual(solve_part_1(lines), 12)

    def test_solve_part_2(self) -> None:
        contents = """""
"abc"
"aaa\\"aaa"
"\\x27"
"""

        lines = contents.split("\n")
        self.assertEqual(solve_part_2(lines), 19)

    def test_answers(self) -> None:
        with open("./2015/08/input.txt", "r") as f:
            lines = f.readlines()
            self.assertEqual(solve_part_1(lines), 1371)
            self.assertEqual(solve_part_2(lines), 2117)


if __name__ == "__main__":
    unittest.main()
