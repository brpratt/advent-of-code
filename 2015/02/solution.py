from collections.abc import Iterable


def parse_dims(s: str) -> tuple[int, int, int]:
    [l, w, h] = s.rstrip().split("x")
    return int(l), int(w), int(h)


def solve_part_1(lines: Iterable[str]) -> int:
    total = 0
    for line in lines:
        l, w, h = parse_dims(line)
        feet = (2 * l * w) + (2 * w * h) + (2 * h * l)
        if l <= h and w <= h:
            feet += l * w
        elif l <= w and h <= w:
            feet += l * h
        else:
            feet += w * h
        total += feet
    return total


def solve_part_2(lines: Iterable[str]) -> int:
    total = 0
    for line in lines:
        l, w, h = parse_dims(line)
        feet = l * w * h
        if l <= h and w <= h:
            feet += l + l + w + w
        elif l <= w and h <= w:
            feet += l + l + h + h
        else:
            feet += w + w + h + h
        total += feet
    return total


if __name__ == "__main__":
    with open("./2015/02/input.txt", "r") as f:
        lines = f.readlines()
        print(solve_part_1(lines))
        print(solve_part_2(lines))
