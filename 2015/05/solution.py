import itertools
import sys
import typing


def has_three_vowels(line: str) -> bool:
    count = 0
    for c in line:
        if c in "aeiou":
            count += 1
    return count >= 3


def has_repeat(line: str) -> bool:
    for pair in itertools.pairwise(line):
        if pair[0] == pair[1]:
            return True
    return False


def has_disallowed(line: str) -> bool:
    disallowed = ["ab", "cd", "pq", "xy"]
    for combo in disallowed:
        if combo in line:
            return True
    return False


def solve_part_1(lines: list[str]) -> int:
    count = 0
    for line in lines:
        if has_three_vowels(line) and has_repeat(line) and not has_disallowed(line):
            count += 1
    return count


def has_double_repeat(line: str) -> bool:
    pair_counts: dict[tuple[str, str], int] = {}
    for pair in itertools.pairwise(line):
        if pair in pair_counts:
            pair_counts[pair] += 1
        else:
            pair_counts[pair] = 1
    for pair, count in pair_counts.items():
        if count == 1:
            continue
        if pair[0] != pair[1]:
            return True
        if line.count(pair[0] + pair[1]) > 1:
            return True
    return False


def has_sandwhich(line: str) -> bool:
    def threes() -> typing.Iterator[tuple[str, str, str]]:
        for idx in range(len(line) - 2):
            yield (line[idx], line[idx + 1], line[idx + 2])

    for three in threes():
        if three[0] == three[2]:
            return True
    return False


def solve_part_2(lines: list[str]) -> int:
    count = 0
    for line in lines:
        if has_double_repeat(line) and has_sandwhich(line):
            count += 1
    return count


if __name__ == "__main__":
    with open("./2015/05/input.txt", "r") as f:
        lines = f.readlines()
        print(solve_part_1(lines))
        print(solve_part_2(lines))
