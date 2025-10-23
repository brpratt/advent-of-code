from typing import NamedTuple


class Point(NamedTuple):
    x: int
    y: int


def move(point: Point, direction: str) -> Point:
    match direction:
        case "<":
            return Point(point.x - 1, point.y)
        case ">":
            return Point(point.x + 1, point.y)
        case "v":
            return Point(point.x, point.y - 1)
        case "^":
            return Point(point.x, point.y + 1)
        case _:
            return point


def solve_part_1(line: str) -> int:
    curr = Point(0, 0)
    points: set[Point] = set()
    points.add(curr)

    for direction in line:
        curr = move(curr, direction)
        points.add(curr)

    return len(points)


def solve_part_2(line: str) -> int:
    orig = Point(0, 0)
    robo = Point(0, 0)
    points: set[Point] = set()
    points.add(orig)

    robo_turn = False

    for direction in line:
        if robo_turn:
            robo = move(robo, direction)
            points.add(robo)
            robo_turn = False
        else:
            orig = move(orig, direction)
            points.add(orig)
            robo_turn = True

    return len(points)


if __name__ == "__main__":
    with open("./2015/03/input.txt", "r") as f:
        line = f.readline().strip()
        print(solve_part_1(line))
        print(solve_part_2(line))
