import unittest

from .solution import solve_part_1, solve_part_2


class TestDay04(unittest.TestCase):
    def test_solve_part_1(self) -> None:
        tests: list[tuple[str, int]] = [
            ("abcdef", 609043),
            ("pqrstuv", 1048970),
        ]

        for dims, expected in tests:
            self.assertEqual(solve_part_1(dims), expected)

    def test_answers(self) -> None:
        with open("./2015/04/input.txt", "r") as f:
            line = f.readline().strip()
            self.assertEqual(solve_part_1(line), 282749)
            self.assertEqual(solve_part_2(line), 9962624)


if __name__ == "__main__":
    unittest.main()
