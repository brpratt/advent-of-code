import unittest

from .solution import solve_part_1, solve_part_2


class TestDay07(unittest.TestCase):
    def test_solve_part_1(self) -> None:
        circuit = """123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i"""

        lines = circuit.split("\n")

        tests: list[tuple[str, int]] = [
            ("d", 72),
            ("e", 507),
            ("f", 492),
            ("g", 114),
            ("h", 65412),
            ("i", 65079),
            ("x", 123),
            ("y", 456),
        ]

        for wire, expected in tests:
            self.assertEqual(solve_part_1(lines, wire), expected)

    def test_answers(self) -> None:
        with open("./2015/07/input.txt", "r") as f:
            lines = f.readlines()
            self.assertEqual(solve_part_1(lines, "a"), 46065)
            self.assertEqual(solve_part_2(lines, "a"), 14134)


if __name__ == "__main__":
    unittest.main()
