import unittest
from collections.abc import Iterable

from .solution import solve_part_1, solve_part_2


class TestDay02(unittest.TestCase):
    def test_solve_part_1(self) -> None:
        tests: list[tuple[Iterable[str], int]] = [
            (["2x3x4\n"], 58),
            (["1x1x10\n"], 43),
            (["2x3x4\n", "1x1x10\n"], 101),
        ]

        for dims, expected in tests:
            self.assertEqual(solve_part_1(dims), expected)

    def test_solve_part_2(self) -> None:
        tests: list[tuple[Iterable[str], int]] = [
            (["2x3x4\n"], 34),
            (["1x1x10\n"], 14),
            (["2x3x4\n", "1x1x10\n"], 48),
        ]

        for dims, expected in tests:
            self.assertEqual(solve_part_2(dims), expected)

    def test_answers(self) -> None:
        with open("2015/02/input.txt", "r") as f:
            lines = f.readlines()
            self.assertEqual(solve_part_1(lines), 1586300)
            self.assertEqual(solve_part_2(lines), 3737498)


if __name__ == "__main__":
    unittest.main()
